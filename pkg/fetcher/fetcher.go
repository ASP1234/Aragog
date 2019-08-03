package fetcher

import (
	"net/url"
)

// Interface exposed for fetching URL
type Fetcher interface {
	fetch(seedUrl url.URL) ([]url.URL, error)
}
