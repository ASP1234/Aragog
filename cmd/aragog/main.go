package main

import (
	module "github.com/ASP1234/Aragog/init"
	"github.com/ASP1234/Aragog/pkg/entity"
	core "github.com/ASP1234/Aragog/pkg/processor"
	"github.com/ASP1234/Aragog/pkg/repository"
	"os"
	"sync"
)

func main() {

	messageQueue := make(chan entity.Message, 100)
	rep := repository.LocalRepository()
	processor := module.Setup(messageQueue, rep)

	crawl(messageQueue, processor)

	processOutput(rep)
}

func crawl(messageQueue chan entity.Message, processor core.Processor) {

	var waitGroup sync.WaitGroup

	for len(messageQueue) > 0 {

		numMsgs := len(messageQueue)
		waitGroup.Add(numMsgs)

		for i := 0; i < numMsgs; i++ {
			go processor.Process(&waitGroup)
		}

		waitGroup.Wait()
	}
}

func processOutput(rep repository.Repository) {

	filePath := "examples/sitemap.txt"

	sitemapFile, err := os.Create(filePath)

	if err != nil {
		panic(err)
	}

	defer sitemapFile.Close()

	data, _, err := rep.BatchScan("")

	if err != nil {
		panic(err)
	}

	for _, page := range data {
		_, err := sitemapFile.WriteString(page.GetAddress().String() + "\n")

		if err != nil {
			panic(err)
		}
	}
}
