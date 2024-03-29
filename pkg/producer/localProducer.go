package producer

import (
	message "github.com/ASP1234/Aragog/pkg/entity"
	customError "github.com/ASP1234/Aragog/pkg/error"
)

// Producer for publishing messages via channels.
type LocalProducer struct {
	messageQueue chan message.Message
}

// Constructs a new LocalProducer object with values being passed as arguments.
func NewLocalProducer(messageQueue chan message.Message) (producer *LocalProducer, err error) {

	if messageQueue == nil {
		err = customError.NewValidationError("messageQueue should not be nil")
		return nil, err
	}

	producer = new(LocalProducer)
	producer.messageQueue = messageQueue

	return producer, nil
}

// Produces the message passed as argument into the channel.
func (producer *LocalProducer) Produce(msg message.Message) (err error) {

	if msg == (message.Message{}) {
		err = customError.NewValidationError("msg should not be nil")
		return err
	}

	producer.messageQueue <- msg

	return err
}
