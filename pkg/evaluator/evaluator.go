// Package evaluator provides use case specific evaluators to filter the links that needs to be fetched further.
package evaluator

import "net/url"

// Interface exposed for evaluating the fetched links based on implementing evaluator rules.
type Evaluator interface {
	Evaluate(seedUrl url.URL, links []*url.URL) (evaluatedLinks []*url.URL, err error)
}
