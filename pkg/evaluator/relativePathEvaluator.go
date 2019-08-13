package evaluator

import (
	customError "github.com/ASP1234/Aragog/pkg/error"
	"net/url"
)

// RelativePathEvaluator to evaluate relative urls and convert them to absolute urls.
type RelativePathEvaluator struct {
}

// Constructs a new RelativePathEvaluator object.
func NewRelativePathEvaluator() (evaluator *RelativePathEvaluator) {
	return new(RelativePathEvaluator)
}

// Evaluate relative urls and convert them to absolute urls with seedUrl as a base.
func (evaluator *RelativePathEvaluator) Evaluate(seedUrl url.URL, links []*url.URL) (
	evaluatedLinks []*url.URL, err error) {

	evaluatedLinks = make([]*url.URL, 0)

	if seedUrl == (url.URL{}) {
		err = customError.NewValidationError("seedUrl should not be empty")

		return evaluatedLinks, err
	}

	for _, link := range links {
		evaluatedLinks = append(evaluatedLinks, seedUrl.ResolveReference(link))
	}

	return evaluatedLinks, err
}
