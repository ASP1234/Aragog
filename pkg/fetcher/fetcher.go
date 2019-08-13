// Package fetcher provides various protocol specific fetchers to fetch data.
package fetcher

import (
	"net/url"
)

// Interface exposed for fetching URL.
type Fetcher interface {
	Fetch(seedUrl url.URL) (links []*url.URL, err error)
}
