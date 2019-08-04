package entity

import (
	customError "Aragog/pkg/error"
	"net/url"
	"time"
)

// WebPage Entity for capturing data related to the fetched web page
type WebPage struct {
	address         url.URL
	links           []url.URL
	lastFetchedDate time.Time
}

// Constructs a new WebPage object with values being passed as arguments
func NewWebPage(address url.URL, links []url.URL, lastFetchedDate time.Time) (webPage *WebPage, err error) {

	if address == (url.URL{}) {
		err = customError.NewValidationError("address should not be empty")
		return nil, err
	}

	if lastFetchedDate == (time.Time{}) {
		err = customError.NewValidationError("lastFetchedDate should not be empty")
		return nil, err
	}

	webPage = new(WebPage)
	webPage.address = address
	webPage.links = links
	webPage.lastFetchedDate = lastFetchedDate

	return webPage, nil
}

// Returns the value of url
func (webPage *WebPage) GetUrl() url.URL {
	return webPage.address
}

// Returns the value of links
func (webPage *WebPage) GetLinks() []url.URL {
	return webPage.links
}

// Returns the value of lastFetchedDate
func (webPage *WebPage) GetLastFetchedDate() time.Time {
	return webPage.lastFetchedDate
}
