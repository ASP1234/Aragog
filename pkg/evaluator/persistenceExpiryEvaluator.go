package evaluator

import (
	"Aragog/pkg/entity"
	"Aragog/pkg/entity/status"
	customError "Aragog/pkg/error"
	"Aragog/pkg/repository"
	"net/url"
	"time"
)

// PersistenceExpiryEvaluator to evaluate expiry of links that are already fetched.
type PersistenceExpiryEvaluator struct {
	expiryTimeOut time.Duration
	rep           repository.Repository
}

// Constructs a new PersistenceExpiryEvaluator object with values being passed as arguments.
func NewPersistenceExpiryEvaluator(expiryTimeOut time.Duration, rep repository.Repository) (
	evaluator *PersistenceExpiryEvaluator, err error) {

	if rep == nil {
		err = customError.NewValidationError("rep should not be nil")
		return nil, err
	}

	evaluator = new(PersistenceExpiryEvaluator)
	evaluator.rep = rep
	evaluator.expiryTimeOut = expiryTimeOut

	return evaluator, err
}

// Evaluate urls based on persistence and configured expiry timeout.
func (evaluator *PersistenceExpiryEvaluator) Evaluate(seedUrl url.URL, links []*url.URL) (
	evaluatedLinks []*url.URL, err error) {

	evaluatedLinks = make([]*url.URL, 0)

	for _, link := range links {

		webPage, err := (evaluator.rep).Get(link.String())

		if err == nil {

			persistedLinkExpiryTime := webPage.GetLastModifiedDate().Add(evaluator.expiryTimeOut)

			if time.Now().After(persistedLinkExpiryTime) {
				evaluatedLinks = append(evaluatedLinks, link)
				evaluator.createInitialEntry(link)
			}

		} else {
			evaluatedLinks = append(evaluatedLinks, link)
			evaluator.createInitialEntry(link)
		}
	}

	return evaluatedLinks, err
}

func (evaluator *PersistenceExpiryEvaluator) createInitialEntry(link *url.URL) {

	firstAttempt := 0
	webPage, err := entity.NewWebPage(link, time.Now(), make([]*url.URL, 0), firstAttempt, status.ToBeFetched)

	if err == nil {
		evaluator.rep.Put(webPage)
	}
}
