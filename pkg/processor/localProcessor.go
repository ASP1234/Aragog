package processor

import (
	"Aragog/pkg/entity"
	"Aragog/pkg/entity/status"
	customError "Aragog/pkg/error"
	"Aragog/pkg/evaluator"
	"Aragog/pkg/fetcher"
	"Aragog/pkg/producer"
	"Aragog/pkg/repository"
	"Aragog/pkg/tools"
	"go.uber.org/zap"
	"net/url"
	"runtime"
	"sync"
	"time"
)

// LocalProcessor for processing messages.
type LocalProcessor struct {
	logger           *zap.SugaredLogger
	maxGoRoutines    int
	maxRetryAttempts int
	messageQueue     chan entity.Message
	fetcher          fetcher.Fetcher
	rep              repository.Repository
	evaluators       []evaluator.Evaluator
	producer         producer.Producer
}

// Constructs a new LocalProcessor object with values being passed as arguments.
func NewLocalProcessor(maxGoRoutines int, maxRetryAttempts int, messageQueue chan entity.Message,
	fetcher fetcher.Fetcher, rep repository.Repository, evaluators []evaluator.Evaluator, producer producer.Producer) (
	localProcessor *LocalProcessor, err error) {

	if maxGoRoutines <= 1 {
		err = customError.NewValidationError("maxGoRoutines should be greater than 1")
		return nil, err
	}

	if maxGoRoutines < 0 {
		err = customError.NewValidationError("maxRetryAttempts should be non-negative")
		return nil, err
	}

	if messageQueue == nil {
		err = customError.NewValidationError("messageQueue should not be nil")
		return nil, err
	}

	if fetcher == nil {
		err = customError.NewValidationError("fetcher should not be nil")
		return nil, err
	}

	if rep == nil {
		err = customError.NewValidationError("rep should not be nil")
		return nil, err
	}

	if producer == nil {
		err = customError.NewValidationError("producer should not be nil")
		return nil, err
	}

	localProcessor = new(LocalProcessor)
	localProcessor.maxGoRoutines = maxGoRoutines
	localProcessor.maxRetryAttempts = maxRetryAttempts
	localProcessor.messageQueue = messageQueue
	localProcessor.fetcher = fetcher
	localProcessor.rep = rep
	localProcessor.evaluators = evaluators
	localProcessor.producer = producer
	localProcessor.logger = tools.GetLogger()

	return localProcessor, err
}

// Process the frontier message (if any) in the messageQueue.
func (processor *LocalProcessor) Process(waitGroup *sync.WaitGroup) {

	select {
	case msg := <-processor.messageQueue:
		processor.process(msg, waitGroup)
	default:
		waitGroup.Done()
		return
	}

	return
}

func (processor *LocalProcessor) process(msg entity.Message, waitGroup *sync.WaitGroup) {

	processor.logger.Infof("Fetching Message: %s", msg.GetLink().String())
	links, err := processor.fetch(msg, waitGroup)

	if err != nil {
		return
	}

	initialAttempt := 0
	webPage, err := entity.NewWebPage(msg.GetLink(), time.Now(), links, initialAttempt, status.Fetched)

	if err != nil {

		switch err.(type) {

		case *customError.ValidationError:
			processor.logger.Errorf("Validation error while constructing webPage for address: %s\n%s", msg.GetLink(), err.Error())

		default:
			processor.logger.Fatalf("Error while constructing webPage for address: %s\n%s", msg.GetLink(), err.Error())
		}

		processor.spawnChildren(0, waitGroup)
		return
	}

	processor.logger.Infof("Storing webPage: %s", webPage.GetAddress().String())
	err = processor.store(webPage, waitGroup)

	if err != nil {
		return
	}

	clonedLinks := append(links[:0:0], links...)

	processor.logger.Infof("Evaluating childLinks for seedUrl: %s", (*msg.GetLink()).String())
	legitChildrenLinks := processor.evaluate(*msg.GetLink(), clonedLinks)

	processor.logger.Infof("Publishing childLinks for seedUrl: %s", (*msg.GetLink()).String())
	processor.publish(legitChildrenLinks)

	processor.spawnChildren(len(legitChildrenLinks), waitGroup)

}

func (processor *LocalProcessor) fetch(msg entity.Message, waitGroup *sync.WaitGroup) (links []*url.URL, err error) {
	links, err = processor.fetcher.Fetch(*msg.GetLink())

	if err != nil {

		switch err.(type) {
		case *customError.DependencyError:
			processor.logger.Errorf("Fetcher failed while fetching seedUrl: %s. \n%s", *msg.GetLink(), err.Error())
			processor.retry(msg)
			processor.spawnChildren(1, waitGroup)

			return links, err
		case *customError.ValidationError:
			processor.logger.Errorf("Validation error while fetching seedUrl: %s. \n%s", *msg.GetLink(), err.Error())
			processor.spawnChildren(0, waitGroup)

			return links, err

		default:
			processor.logger.Fatalf("Error while fetching seedUrl: %s. \n%s", *msg.GetLink(), err.Error())
			processor.spawnChildren(0, waitGroup)

			return links, err
		}
	}

	return links, err
}

func (processor *LocalProcessor) store(page *entity.WebPage, waitGroup *sync.WaitGroup) (err error) {
	_, err = processor.rep.Put(page)

	if err != nil {

		switch err.(type) {
		case *customError.DependencyError:
			processor.logger.Errorf("Put failed while storing id: %s. \n%s", page.GetAddress().String(), err.Error())
			msg, _ := entity.NewMessage(page.GetAddress())
			processor.retry(*msg)
			processor.spawnChildren(1, waitGroup)

			return err
		case *customError.ValidationError:
			processor.logger.Errorf("Validation error while storing id: %s. \n%s", page.GetAddress().String(), err.Error())
			processor.spawnChildren(0, waitGroup)

			return err

		default:
			processor.logger.Fatalf("Error while storing id: %s. \n%s", page.GetAddress().String(), err.Error())
			processor.spawnChildren(0, waitGroup)

			return err
		}
	}

	return err
}

func (processor *LocalProcessor) retry(msg entity.Message) {

	webPage, err := processor.rep.Get(msg.GetLink().String())
	exhaustedRetryAttempts := 0

	if err == nil {
		exhaustedRetryAttempts = webPage.GetRetryAttempts()
		webPage, err = entity.NewWebPage(msg.GetLink(), webPage.GetLastModifiedDate(), webPage.GetLinks(), exhaustedRetryAttempts+1, status.ToBeFetched)
	} else {
		webPage, err = entity.NewWebPage(msg.GetLink(), time.Now(), make([]*url.URL, 0), exhaustedRetryAttempts, status.ToBeFetched)
	}

	if err != nil {
		switch err.(type) {
		case *customError.ValidationError:
			processor.logger.Errorf("Validation error while constructing webPage for address: %s\n%s", msg.GetLink(), err.Error())
		default:
			processor.logger.Fatalf("Error while constructing webPage for address: %s\n%s", msg.GetLink(), err.Error())
		}

		return

	} else {
		if exhaustedRetryAttempts < processor.maxRetryAttempts {
			processor.rep.Put(webPage)
			err = processor.producer.Produce(msg)

			if err != nil {
				processor.logger.Fatal("Error while producing msg: %s \n%s", msg.GetLink().String(), err.Error())
			}
		}
	}
}

func (processor *LocalProcessor) evaluate(seedUrl url.URL, clonedLinks []*url.URL) []*url.URL {

	for _, filter := range processor.evaluators {
		filteredLinks, err := filter.Evaluate(seedUrl, clonedLinks)

		if err != nil {
			processor.logger.Error("Error while using filter. \n%s", err.Error())
		} else {
			clonedLinks = filteredLinks
		}
	}

	return clonedLinks
}

func (processor *LocalProcessor) publish(links []*url.URL) {

	for _, link := range links {
		msgToPublish, err := entity.NewMessage(link)

		if err != nil {
			processor.logger.Error("Error while constructing message for link: %s. \n%s", link.String(), err.Error())
		} else {
			err = processor.producer.Produce(*msgToPublish)

			if err != nil {
				processor.logger.Fatal("Error while producing msg: %s \n%s", link.String(), err.Error())
			}
		}
	}
}

func (processor *LocalProcessor) spawnChildren(numChildren int, waitGroup *sync.WaitGroup) {

	remainingWorkers := processor.maxGoRoutines - runtime.NumGoroutine()

	if remainingWorkers < numChildren {
		numChildren = remainingWorkers
	}

	waitGroup.Add(numChildren)
	waitGroup.Done()

	processor.logger.Infof("Spawning %d children", numChildren)

	for i := 0; i < numChildren; i++ {
		go processor.Process(waitGroup)
	}
}
