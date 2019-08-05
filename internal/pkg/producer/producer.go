package producer

import (
	message "Aragog/internal/pkg/entity"
)

// Interface exposed for producing message
type Producer interface {
	Produce(msg message.Message) (err error)
}
