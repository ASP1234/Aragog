package fetcher

import (
	customError "github.com/ASP1234/Aragog/pkg/error"
	"golang.org/x/net/html"
	"net/http"
	"net/url"
)

// HTTP protocol implementation of fetcher for fetching URL.
type HttpFetcher struct {
}

// Constructs a new HttpFetcher object.
func NewHttpFetcher() (httpFetcher *HttpFetcher, err error) {
	return new(HttpFetcher), nil
}

// Fetches the seedUrl.
// Returns the slice of URLs present in the response body of seedUrl.
func (httpFetcher *HttpFetcher) Fetch(seedUrl url.URL) (links []*url.URL, err error) {

	if seedUrl == (url.URL{}) {
		err = customError.NewValidationError("seedUrl should not be empty")

		return nil, err
	}

	resp, err := http.Get(seedUrl.String())

	if err != nil {
		err = customError.NewDependencyError(err.Error())

		return nil, err
	}

	return parse(resp), nil
}

func parse(resp *http.Response) (links []*url.URL) {

	body := resp.Body
	tokenizer := html.NewTokenizer(body)

	for tokenizer.Next() != html.ErrorToken {
		tokenType := tokenizer.Next()

		if tokenType == html.StartTagToken {
			token := tokenizer.Token()

			if isAnchor(token) {
				link, err := getChildLink(token)

				if err == nil && link != nil {
					links = append(links, link)
				}
			}
		}
	}

	return links
}

func isAnchor(token html.Token) bool {

	const anchor = "a"

	return token.Data == anchor
}

func getChildLink(token html.Token) (link *url.URL, err error) {

	const hrefKey = "href"

	for _, attr := range token.Attr {
		if attr.Key == hrefKey {
			link, err = url.Parse(attr.Val)
		}
	}

	return link, err
}
