package entity

import (
	"Aragog/pkg/entity/status"
	"net/url"
	"reflect"
	"testing"
	"time"
)

const (
	fetchedUrlString     = "https://monzo.com/"
	InvalidRetryAttempts = -1
	InvalidStatus        = ""
	ValidRetryAttempt    = 0
)

func TestNewWebPage(t *testing.T) {

	type args struct {
		address          *url.URL
		lastModifiedDate time.Time
		links            []*url.URL
		retryAttempts    int
		status           string
	}

	tests := []struct {
		name        string
		args        args
		wantWebPage *WebPage
		wantErr     bool
	}{
		0: {name: "ValidArguments",
			args:        args{address: fetchedURL(), lastModifiedDate: sampleLastModifiedDate(), links: nil, retryAttempts: ValidRetryAttempt, status: status.Fetched},
			wantWebPage: sampleWebPage(),
			wantErr:     false},
		1: {name: "InvalidAddress",
			args:        args{address: nil, lastModifiedDate: sampleLastModifiedDate(), links: nil, retryAttempts: ValidRetryAttempt, status: status.ToBeFetched},
			wantWebPage: nil,
			wantErr:     true},
		2: {name: "InvalidLastModifiedDate",
			args:        args{address: fetchedURL(), lastModifiedDate: time.Time{}, links: nil, retryAttempts: ValidRetryAttempt, status: status.Fetched},
			wantWebPage: nil,
			wantErr:     true},
		3: {name: "InvalidRetryAttempts",
			args:        args{address: fetchedURL(), lastModifiedDate: sampleLastModifiedDate(), links: nil, retryAttempts: InvalidRetryAttempts, status: status.Fetched},
			wantWebPage: nil,
			wantErr:     true},
		4: {name: "InvalidStatus",
			args:        args{address: fetchedURL(), lastModifiedDate: sampleLastModifiedDate(), links: nil, retryAttempts: ValidRetryAttempt, status: InvalidStatus},
			wantWebPage: nil,
			wantErr:     true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotWebPage, err := NewWebPage(tt.args.address, tt.args.lastModifiedDate, tt.args.links, tt.args.retryAttempts, tt.args.status)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewWebPage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotWebPage, tt.wantWebPage) {
				t.Errorf("NewWebPage() gotWebPage = %v, want %v", gotWebPage, tt.wantWebPage)
			}
		})
	}
}

func TestWebPage_GetAddress(t *testing.T) {

	type fields struct {
		address          *url.URL
		lastModifiedDate time.Time
		links            []*url.URL
		retryAttempts    int
		status           string
	}

	tests := []struct {
		name   string
		fields fields
		want   *url.URL
	}{
		0: {name: "ValidFields",
			fields: fields{address: fetchedURL(), lastModifiedDate: sampleLastModifiedDate(), links: nil, retryAttempts: ValidRetryAttempt, status: status.Fetched},
			want:   fetchedURL()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			webPage := &WebPage{
				address:          tt.fields.address,
				lastModifiedDate: tt.fields.lastModifiedDate,
				links:            tt.fields.links,
				retryAttempts:    tt.fields.retryAttempts,
				status:           tt.fields.status,
			}
			if got := webPage.GetAddress(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWebPage_GetLastModifiedDate(t *testing.T) {

	type fields struct {
		address          *url.URL
		lastModifiedDate time.Time
		links            []*url.URL
		retryAttempts    int
		status           string
	}

	tests := []struct {
		name   string
		fields fields
		want   time.Time
	}{
		0: {name: "ValidFields",
			fields: fields{address: fetchedURL(), lastModifiedDate: sampleLastModifiedDate(), links: nil, retryAttempts: ValidRetryAttempt, status: status.Fetched},
			want:   sampleLastModifiedDate()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			webPage := &WebPage{
				address:          tt.fields.address,
				lastModifiedDate: tt.fields.lastModifiedDate,
				links:            tt.fields.links,
				retryAttempts:    tt.fields.retryAttempts,
				status:           tt.fields.status,
			}
			if got := webPage.GetLastModifiedDate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLastModifiedDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWebPage_GetLinks(t *testing.T) {

	type fields struct {
		address          *url.URL
		lastModifiedDate time.Time
		links            []*url.URL
		retryAttempts    int
		status           string
	}

	tests := []struct {
		name   string
		fields fields
		want   []*url.URL
	}{
		0: {name: "ValidFields",
			fields: fields{address: fetchedURL(), lastModifiedDate: sampleLastModifiedDate(), links: nil, retryAttempts: ValidRetryAttempt, status: status.Fetched},
			want:   nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			webPage := &WebPage{
				address:          tt.fields.address,
				lastModifiedDate: tt.fields.lastModifiedDate,
				links:            tt.fields.links,
				retryAttempts:    tt.fields.retryAttempts,
				status:           tt.fields.status,
			}
			if got := webPage.GetLinks(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLinks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWebPage_GetRetryAttempts(t *testing.T) {

	type fields struct {
		address          *url.URL
		lastModifiedDate time.Time
		links            []*url.URL
		retryAttempts    int
		status           string
	}

	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		0: {name: "ValidFields",
			fields: fields{address: fetchedURL(), lastModifiedDate: sampleLastModifiedDate(), links: nil, retryAttempts: ValidRetryAttempt, status: status.Fetched},
			want:   0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			webPage := &WebPage{
				address:          tt.fields.address,
				lastModifiedDate: tt.fields.lastModifiedDate,
				links:            tt.fields.links,
				retryAttempts:    tt.fields.retryAttempts,
				status:           tt.fields.status,
			}
			if got := webPage.GetRetryAttempts(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWebPage_GetStatus(t *testing.T) {

	type fields struct {
		address          *url.URL
		lastModifiedDate time.Time
		links            []*url.URL
		retryAttempts    int
		status           string
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		0: {name: "ValidFields",
			fields: fields{address: fetchedURL(), lastModifiedDate: sampleLastModifiedDate(), links: nil, retryAttempts: ValidRetryAttempt, status: status.Fetched},
			want:   status.Fetched},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			webPage := &WebPage{
				address:          tt.fields.address,
				lastModifiedDate: tt.fields.lastModifiedDate,
				links:            tt.fields.links,
				retryAttempts:    tt.fields.retryAttempts,
				status:           tt.fields.status,
			}
			if got := webPage.GetStatus(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func fetchedURL() (link *url.URL) {

	urlPtr, _ := url.Parse(fetchedUrlString)
	link = urlPtr

	return link
}

func sampleLastModifiedDate() time.Time {

	sampleTime, _ := time.Parse(time.Now().Format(time.UnixDate), time.UnixDate)
	sampleTime = sampleTime.AddDate(1, 1, 1)

	return sampleTime
}

func sampleWebPage() *WebPage {

	webPage, _ := NewWebPage(fetchedURL(), sampleLastModifiedDate(), nil, ValidRetryAttempt, status.Fetched)

	return webPage
}
