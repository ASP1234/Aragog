// Package producer provides various custom implementations of producer to produce the message.
package producer

import (
	message "github.com/ASP1234/Aragog/pkg/entity"
)

// Interface exposed for producing message.
type Producer interface {
	Produce(msg message.Message) (err error)
}
