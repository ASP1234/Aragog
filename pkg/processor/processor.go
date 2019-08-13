// Package processor provides various custom implementations of processor to process the message.
package processor

import "sync"

// Interface exposed for processing message.
type Processor interface {
	Process(waitGroup *sync.WaitGroup)
}
