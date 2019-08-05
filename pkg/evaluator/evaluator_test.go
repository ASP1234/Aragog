package evaluator

import "net/url"

func createURL(rawString string) (link *url.URL) {

	urlPtr, _ := url.Parse(rawString)
	link = urlPtr

	return link
}
