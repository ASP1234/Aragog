// Package entity provides various entities to capture the business model.
package entity

import (
	customError "github.com/ASP1234/Aragog/pkg/error"
	"net/url"
)

// Message Entity for capturing data related to the address to be fetched.
type Message struct {
	link *url.URL
}

// Constructs a new Message object with values being passed as arguments.
func NewMessage(link *url.URL) (msg *Message, err error) {

	if link == nil {
		err = customError.NewValidationError("link should not be nil")
		return nil, err
	}

	msg = new(Message)
	msg.link = link

	return msg, err
}

// Returns the value of link.
func (msg *Message) GetLink() *url.URL {

	return msg.link
}
