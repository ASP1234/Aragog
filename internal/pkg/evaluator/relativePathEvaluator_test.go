package evaluator

import (
	"net/url"
	"reflect"
	"testing"
)

const (
	absoluteRelativeUrlString = "https://monzo.com/about/"
	absoluteUrlString         = "https://monzo.com/"
	relativeUrlString         = "/about/"
)

func TestNewRelativePathEvaluator(t *testing.T) {
	tests := []struct {
		name          string
		wantEvaluator *RelativePathEvaluator
	}{
		0: {name: "Constructor", wantEvaluator: NewRelativePathEvaluator()}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotEvaluator := NewRelativePathEvaluator(); !reflect.DeepEqual(gotEvaluator, tt.wantEvaluator) {
				t.Errorf("NewRelativePathEvaluator() = %v, want %v", gotEvaluator, tt.wantEvaluator)
			}
		})
	}
}

func TestRelativePathEvaluator_Evaluate(t *testing.T) {

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
			args:               args{seedUrl: url.URL{}, links: append(links, createURL(relativeUrlString), createURL(absoluteUrlString))},
			wantEvaluatedLinks: links,
			wantErr:            true},
		1: {name: "EmptyLinks",
			args:               args{seedUrl: *createURL(absoluteUrlString), links: links},
			wantEvaluatedLinks: links,
			wantErr:            false},
		2: {name: "ValidArguments",
			args:               args{seedUrl: *createURL(absoluteUrlString), links: append(links, createURL(relativeUrlString), createURL(absoluteUrlString))},
			wantEvaluatedLinks: append(links, createURL(absoluteRelativeUrlString), createURL(absoluteUrlString)),
			wantErr:            false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			evaluator := &RelativePathEvaluator{}
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
