package evaluator

import (
	"Aragog/internal/pkg/repository"
	customError "Aragog/pkg/error"
	"net/url"
	"time"
)

// PersistenceExpiryEvaluator to evaluate expiry of links that are already fetched
type PersistenceExpiryEvaluator struct {
	rep           *repository.Repository
	expiryTimeOut time.Duration
}

// Constructs a new PersistenceExpiryEvaluator object with values being passed as arguments
func NewPersistenceExpiryEvaluator(rep *repository.Repository, expiryTimeOut time.Duration) (evaluator *PersistenceExpiryEvaluator, err error) {

	if rep == nil {
		err = customError.NewValidationError("rep should not be nil")
		return nil, err
	}

	evaluator = new(PersistenceExpiryEvaluator)
	evaluator.rep = rep
	evaluator.expiryTimeOut = expiryTimeOut

	return evaluator, err
}

// Evaluate urls based on persistence and configured expiry timeout
func (evaluator *PersistenceExpiryEvaluator) Evaluate(seedUrl url.URL, links []*url.URL) (evaluatedLinks []*url.URL, err error) {

	evaluatedLinks = make([]*url.URL, 0)

	for _, link := range links {

		webPage, err := (*evaluator.rep).Get(link.String())

		if err == nil {

			persistedLinkExpiryTime := webPage.GetLastFetchedDate().Add(evaluator.expiryTimeOut)

			if time.Now().After(persistedLinkExpiryTime) {
				evaluatedLinks = append(evaluatedLinks, link)
			}
		} else {
			evaluatedLinks = append(evaluatedLinks, link)
		}
	}

	return evaluatedLinks, err
}
