package evaluator

import (
	"net/url"
	"reflect"
	"testing"
)

const (
	externalLinkUrlString = "https://facebook.com"
	seedLinkUrlString     = "https://monzo.com/about/"
	seedUrlString         = "https://monzo.com/"
)

func TestNewSubDomainEvaluator(t *testing.T) {

	tests := []struct {
		name          string
		wantEvaluator *SubDomainEvaluator
	}{
		0: {name: "Constructor", wantEvaluator: NewSubDomainEvaluator()}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotEvaluator := NewSubDomainEvaluator(); !reflect.DeepEqual(gotEvaluator, tt.wantEvaluator) {
				t.Errorf("NewSubDomainEvaluator() = %v, want %v", gotEvaluator, tt.wantEvaluator)
			}
		})
	}
}

func TestSubDomainEvaluator_Evaluate(t *testing.T) {

	links := make([]*url.URL, 0)

	type args struct {
		seedUrl url.URL
		links   []*url.URL
	}

	tests := []struct {
		name               string
		args               args
		wantEvaluatedLinks []*url.URL
		wantErr            bool
	}{
		0: {name: "EmptySeedUrl",
			args:               args{seedUrl: url.URL{}, links: append(links, createURL(seedLinkUrlString), createURL(externalLinkUrlString))},
			wantEvaluatedLinks: links,
			wantErr:            true},
		1: {name: "EmptyLinks",
			args:               args{seedUrl: *createURL(seedUrlString), links: links},
			wantEvaluatedLinks: links,
			wantErr:            false},
		2: {name: "ValidArguments",
			args:               args{seedUrl: *createURL(seedUrlString), links: append(links, createURL(seedLinkUrlString), createURL(externalLinkUrlString))},
			wantEvaluatedLinks: append(links, createURL(seedLinkUrlString)),
			wantErr:            false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			evaluator := &SubDomainEvaluator{}
			gotEvaluatedLinks, err := evaluator.Evaluate(tt.args.seedUrl, tt.args.links)
			if (err != nil) != tt.wantErr {
				t.Errorf("Evaluate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotEvaluatedLinks, tt.wantEvaluatedLinks) {
				t.Errorf("Evaluate() gotEvaluatedLinks = %v, want %v", gotEvaluatedLinks, tt.wantEvaluatedLinks)
			}
		})
	}
}
