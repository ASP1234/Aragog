package repository

import (
	"Aragog/pkg/entity"
	"Aragog/pkg/entity/status"
	"net/url"
	"reflect"
	"sync"
	"testing"
	"time"
)

const (
	fetchedUrlString   = "https://monzo.com/"
	validRetryAttempts = 0
)

func TestLocalRepository(t *testing.T) {

	repository := LocalRepository()

	tests := []struct {
		name string
		want *localRepository
	}{
		0: {"Singleton", repository},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LocalRepository(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LocalRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_localRepository_Put(t *testing.T) {

	webPage := sampleWebPage()
	emptyRepo := make(map[string]*entity.WebPage)

	repo := make(map[string]*entity.WebPage)
	repo[webPage.GetAddress().String()] = webPage

	type fields struct {
		repo map[string]*entity.WebPage
		mu   sync.RWMutex
	}
	type args struct {
		webPage *entity.WebPage
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantId  string
		wantErr bool
	}{
		0: {name: "Create",
			fields:  fields{repo: emptyRepo, mu: sync.RWMutex{}},
			args:    args{webPage: webPage},
			wantId:  fetchedUrlString,
			wantErr: false},
		1: {name: "Update",
			fields:  fields{repo: repo, mu: sync.RWMutex{}},
			args:    args{webPage: webPage},
			wantId:  fetchedUrlString,
			wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rep := &localRepository{
				repo: tt.fields.repo,
				mu:   tt.fields.mu,
			}
			gotId, err := rep.Put(tt.args.webPage)
			if (err != nil) != tt.wantErr {
				t.Errorf("Put() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotId != tt.wantId {
				t.Errorf("Put() gotId = %v, want %v", gotId, tt.wantId)
			}
		})
	}
}

func Test_localRepository_Get(t *testing.T) {

	webPage := sampleWebPage()
	emptyRepo := make(map[string]*entity.WebPage)

	repo := make(map[string]*entity.WebPage)
	repo[webPage.GetAddress().String()] = webPage

	type fields struct {
		repo map[string]*entity.WebPage
		mu   sync.RWMutex
	}

	type args struct {
		id string
	}

	tests := []struct {
		name        string
		fields      fields
		args        args
		wantWebPage *entity.WebPage
		wantErr     bool
	}{
		0: {name: "WebPagePresent",
			fields:      fields{repo: repo, mu: sync.RWMutex{}},
			args:        args{id: fetchedUrlString},
			wantWebPage: webPage,
			wantErr:     false},

		1: {name: "NoWebPagePresent",
			fields:      fields{repo: emptyRepo, mu: sync.RWMutex{}},
			args:        args{id: fetchedUrlString},
			wantWebPage: nil,
			wantErr:     true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rep := &localRepository{
				repo: tt.fields.repo,
				mu:   tt.fields.mu,
			}
			gotWebPage, err := rep.Get(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotWebPage, tt.wantWebPage) {
				t.Errorf("Get() gotWebPage = %v, want %v", gotWebPage, tt.wantWebPage)
			}
		})
	}
}

func Test_localRepository_BatchScan(t *testing.T) {

	webPage := sampleWebPage()
	emptyRepo := make(map[string]*entity.WebPage)

	repo := make(map[string]*entity.WebPage)
	repo[webPage.GetAddress().String()] = webPage

	type fields struct {
		repo map[string]*entity.WebPage
		mu   sync.RWMutex
	}

	type args struct {
		exclusiveStartKey string
	}

	tests := []struct {
		name            string
		fields          fields
		args            args
		wantScanResults []*entity.WebPage
		wantErr         bool
	}{
		0: {name: "WebPagesPresent",
			fields:          fields{repo: repo, mu: sync.RWMutex{}},
			args:            *new(args),
			wantScanResults: append(make([]*entity.WebPage, 0), webPage),
			wantErr:         false},

		1: {name: "NoWebPagesPresent",
			fields:          fields{repo: emptyRepo, mu: sync.RWMutex{}},
			args:            *new(args),
			wantScanResults: make([]*entity.WebPage, 0),
			wantErr:         false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rep := &localRepository{
				repo: tt.fields.repo,
				mu:   tt.fields.mu,
			}
			gotScanResults, _, err := rep.BatchScan(tt.args.exclusiveStartKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("BatchScan() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotScanResults, tt.wantScanResults) {
				t.Errorf("BatchScan() gotScanResults = %v, want %v", gotScanResults, tt.wantScanResults)
			}
		})
	}
}

func fetchedURL() (link *url.URL) {

	urlPtr, _ := url.Parse(fetchedUrlString)
	link = urlPtr

	return link
}

func sampleLastFetchedDate() time.Time {

	sampleTime, _ := time.Parse(time.Now().Format(time.UnixDate), time.UnixDate)
	sampleTime = sampleTime.AddDate(1, 1, 1)

	return sampleTime
}

func sampleWebPage() *entity.WebPage {

	webPage, _ := entity.NewWebPage(fetchedURL(), sampleLastFetchedDate(), nil, validRetryAttempts, status.ToBeFetched)

	return webPage
}
