package processor

import "sync"

// Interface exposed for processing message
type Processor interface {
	Process(waitGroup *sync.WaitGroup)
}
