package message

import (
	customError "Aragog/pkg/error"
	"net/url"
)

// Message Entity for capturing data related to the url to be fetched
type Message struct {
	link url.URL
}

// Constructs a new Message object with values being passed as arguments
func New(link url.URL) (msg *Message, err error) {

	if link == (url.URL{}) {
		err = customError.NewValidationError("link should not be empty")
		return nil, err
	}

	msg = new(Message)
	msg.link = link

	return msg, err
}

// Returns the value of link
func (msg *Message) GetLink() url.URL {

	return msg.link
}
