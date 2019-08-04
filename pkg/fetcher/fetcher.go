package fetcher

import (
	"net/url"
)

// Interface exposed for fetching URL
type Fetcher interface {
	Fetch(seedUrl url.URL) (links []*url.URL, err error)
}
