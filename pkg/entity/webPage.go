package entity

import (
	customError "github.com/ASP1234/Aragog/pkg/error"
	"net/url"
	"time"
)

// WebPage Entity for capturing data related to the fetched web page.
type WebPage struct {
	address          *url.URL
	lastModifiedDate time.Time
	links            []*url.URL
	retryAttempts    int
	status           string
}

// Constructs a new WebPage object with values being passed as arguments.
func NewWebPage(address *url.URL, lastModifiedDate time.Time,
	links []*url.URL, retryAttempts int, status string) (webPage *WebPage, err error) {

	if address == nil {
		err = customError.NewValidationError("address should not be nil")
		return nil, err
	}

	if lastModifiedDate == (time.Time{}) {
		err = customError.NewValidationError("lastModifiedDate should not be empty")
		return nil, err
	}

	if retryAttempts < 0 {
		err = customError.NewValidationError("retryAttempts should not be negative")
		return nil, err
	}

	if len(status) == 0 {
		err = customError.NewValidationError("status should not be empty")
		return nil, err
	}

	webPage = new(WebPage)
	webPage.address = address
	webPage.lastModifiedDate = lastModifiedDate
	webPage.links = links
	webPage.retryAttempts = retryAttempts
	webPage.status = status

	return webPage, nil
}

// Returns the value of address.
func (webPage *WebPage) GetAddress() *url.URL {
	return webPage.address
}

// Returns the value of lastModifiedDate.
func (webPage *WebPage) GetLastModifiedDate() time.Time {
	return webPage.lastModifiedDate
}

// Returns the value of links.
func (webPage *WebPage) GetLinks() []*url.URL {
	return webPage.links
}

// Returns the count of retry attempts.
func (webPage *WebPage) GetRetryAttempts() int {
	return webPage.retryAttempts
}

// Returns the value of status.
func (webPage *WebPage) GetStatus() string {
	return webPage.status
}
