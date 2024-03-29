package evaluator

import (
	"github.com/ASP1234/Aragog/pkg/entity"
	"github.com/ASP1234/Aragog/pkg/entity/status"
	"github.com/ASP1234/Aragog/pkg/repository"
	"net/url"
	"reflect"
	"testing"
	"time"
)

const (
	currentSeedUrlString  = "https://monzo.com/"
	notPersistedUrlString = "https://monzo.com/about/"
	persistedUrlString    = "https://monzo.com/careers/"
	sampleRetryAttempt    = 1
)

func TestNewPersistenceExpiryEvaluator(t *testing.T) {

	var rep repository.Repository = repository.LocalRepository()
	evaluator, _ := NewPersistenceExpiryEvaluator(time.Duration(1*time.Microsecond), rep)

	type args struct {
		expiryTimeOut time.Duration
		rep           repository.Repository
	}

	tests := []struct {
		name          string
		args          args
		wantEvaluator *PersistenceExpiryEvaluator
		wantErr       bool
	}{
		0: {name: "ValidArguments",
			args:          args{expiryTimeOut: time.Microsecond, rep: rep},
			wantEvaluator: evaluator,
			wantErr:       false},
		1: {name: "NilRepository",
			args:          args{expiryTimeOut: time.Microsecond, rep: nil},
			wantEvaluator: nil,
			wantErr:       true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotEvaluator, err := NewPersistenceExpiryEvaluator(tt.args.expiryTimeOut, tt.args.rep)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPersistenceExpiryEvaluator() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotEvaluator, tt.wantEvaluator) {
				t.Errorf("NewPersistenceExpiryEvaluator() gotEvaluator = %v, want %v", gotEvaluator, tt.wantEvaluator)
			}
		})
	}
}

func TestPersistenceExpiryEvaluator_Evaluate(t *testing.T) {

	var rep repository.Repository = repository.LocalRepository()
	links := make([]*url.URL, 0)
	webPage, _ := entity.NewWebPage(createURL(persistedUrlString), time.Now(), links, sampleRetryAttempt, status.Fetched)
	rep.Put(webPage)

	type fields struct {
		rep           repository.Repository
		expiryTimeOut time.Duration
	}

	type args struct {
		seedUrl url.URL
		links   []*url.URL
	}

	tests := []struct {
		name               string
		fields             fields
		args               args
		wantEvaluatedLinks []*url.URL
		wantErr            bool
	}{
		0: {name: "LinkNotPersisted",
			fields:             fields{rep: rep, expiryTimeOut: time.Duration(1 * time.Microsecond)},
			args:               args{seedUrl: *createURL(currentSeedUrlString), links: append(links, createURL(notPersistedUrlString))},
			wantEvaluatedLinks: append(links, createURL(notPersistedUrlString)),
			wantErr:            false},
		1: {name: "PersistedLinkExpired",
			fields:             fields{rep: rep, expiryTimeOut: time.Duration(1 * time.Microsecond)},
			args:               args{seedUrl: *createURL(currentSeedUrlString), links: append(links, createURL(persistedUrlString))},
			wantEvaluatedLinks: append(links, createURL(persistedUrlString)),
			wantErr:            false},
		2: {name: "PersistedLinkNotExpired",
			fields:             fields{rep: rep, expiryTimeOut: time.Duration(1 * time.Hour)},
			args:               args{seedUrl: *createURL(currentSeedUrlString), links: append(links, createURL(persistedUrlString))},
			wantEvaluatedLinks: links,
			wantErr:            false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			evaluator := &PersistenceExpiryEvaluator{
				rep:           tt.fields.rep,
				expiryTimeOut: tt.fields.expiryTimeOut,
			}
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
