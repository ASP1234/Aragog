package evaluator

import (
	customError "Aragog/pkg/error"
	"net/url"
)

// SubDomainEvaluator to evaluate urls based on domain.
type SubDomainEvaluator struct {
}

// Constructs a new SubDomainEvaluator object.
func NewSubDomainEvaluator() (evaluator *SubDomainEvaluator) {
	return new(SubDomainEvaluator)
}

// Evaluate urls based on seedUrl domain.
func (evaluator *SubDomainEvaluator) Evaluate(seedUrl url.URL, links []*url.URL) (evaluatedLinks []*url.URL, err error) {

	evaluatedLinks = make([]*url.URL, 0)

	if seedUrl == (url.URL{}) {
		err = customError.NewValidationError("seedUrl should not be empty")

		return evaluatedLinks, err
	}

	for _, link := range links {

		if seedUrl.Hostname() == link.Hostname() {
			evaluatedLinks = append(evaluatedLinks, seedUrl.ResolveReference(link))
		}
	}

	return evaluatedLinks, err
}
