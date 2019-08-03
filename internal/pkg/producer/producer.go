package producer

import (
	message "Aragog/internal/pkg/entity"
)

// Interface exposed for producing message
type Producer interface {
	produce(msg message.Message)
}
