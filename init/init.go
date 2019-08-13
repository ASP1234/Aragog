package init

import (
	"Aragog/configs"
	"Aragog/internal/pkg/entity"
	core "Aragog/internal/pkg/processor"
	"Aragog/internal/pkg/producer"
	"Aragog/internal/pkg/repository"
	"Aragog/pkg/evaluator"
	"Aragog/pkg/fetcher"
	"gopkg.in/yaml.v2"
	"net/url"
	"os"
	"runtime"
	"time"
)

// Initializes crawler based on the passed configuration
func Setup(messageQueue chan entity.Message, rep repository.Repository) (processor core.Processor) {

	cfg := getEnvConfiguration()

	runtime.GOMAXPROCS(cfg.OS.MaxProcessors)

	producer := getProducer(messageQueue)
	plantSeedUrl(cfg.Producer.SeedUrl, producer)

	fetcher := getFetcher()
	evaluators := getEvaluators(cfg, rep)

	processor, err := core.NewLocalProcessor(cfg.Processor.MaxRoutines, cfg.Processor.MaxRetryAttempts, messageQueue,
		fetcher, rep, evaluators, producer)

	if err != nil {
		panic(err)
	}
	return processor
}

func getEnvConfiguration() (cfg configs.Init) {

	cfgRelativeFilePath := "configs/config.yml"
	cfgFile, err := os.Open(cfgRelativeFilePath)

	if err != nil {
		panic(err)
	}

	decoder := yaml.NewDecoder(cfgFile)
	err = decoder.Decode(&cfg)
	if err != nil {
		panic(err)
	}

	return cfg
}

func getProducer(messageQueue chan entity.Message) (p producer.Producer) {

	p, err := producer.NewLocalProducer(messageQueue)

	if err != nil {
		panic(err)
	}

	return p
}

func plantSeedUrl(seedUrl string, producer producer.Producer) {

	link, err := url.Parse(seedUrl)

	if err != nil {
		panic(err)
	}

	msg, err := entity.NewMessage(link)

	if err != nil {
		panic(err)
	}

	err = producer.Produce(*msg)

	if err != nil {
		panic(err)
	}
}

func getFetcher() (f fetcher.Fetcher) {

	f, err := fetcher.NewHttpFetcher()

	if err != nil {
		panic(err)
	}

	return f
}

func getEvaluators(cfg configs.Init, rep repository.Repository) (evaluators []evaluator.Evaluator) {

	evaluators = make([]evaluator.Evaluator, 0)
	evaluators = append(evaluators, evaluator.NewRelativePathEvaluator())
	evaluators = append(evaluators, evaluator.NewSubDomainEvaluator())

	eval, err := evaluator.NewPersistenceExpiryEvaluator(time.Hour, rep)

	if err != nil {
		panic(err)
	}

	evaluators = append(evaluators, eval)

	return evaluators
}

