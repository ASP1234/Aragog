package entity

import (
	"net/url"
	"reflect"
	"testing"
	"time"
)

const fetchedUrlString = "https://monzo.com/"

func TestNewWebPage(t *testing.T) {
	type args struct {
		address         url.URL
		links           []url.URL
		lastFetchedDate time.Time
	}
	tests := []struct {
		name        string
		args        args
		wantWebPage *WebPage
		wantErr     bool
	}{
		0: {name: "ValidArguments",
			args:        args{address: fetchedURL(), links: nil, lastFetchedDate: sampleLastFetchedDate()},
			wantWebPage: sampleWebPage(),
			wantErr:     false},
		1: {name: "InvalidAddress",
			args:        args{address: url.URL{}, links: nil, lastFetchedDate: sampleLastFetchedDate()},
			wantWebPage: nil,
			wantErr:     true},
		2: {name: "InvalidLastFetchedDate",
			args:        args{address: fetchedURL(), links: nil, lastFetchedDate: time.Time{}},
			wantWebPage: nil,
			wantErr:     true}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotWebPage, err := NewWebPage(tt.args.address, tt.args.links, tt.args.lastFetchedDate)
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

func TestWebPage_GetLastFetchedDate(t *testing.T) {
	type fields struct {
		address         url.URL
		links           []url.URL
		lastFetchedDate time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   time.Time
	}{
		0: {name: "ValidFields",
			fields: fields{address: fetchedURL(), links: nil, lastFetchedDate: sampleLastFetchedDate()},
			want:   sampleLastFetchedDate()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			webPage := &WebPage{
				address:         tt.fields.address,
				links:           tt.fields.links,
				lastFetchedDate: tt.fields.lastFetchedDate,
			}
			if got := webPage.GetLastFetchedDate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLastFetchedDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWebPage_GetLinks(t *testing.T) {
	type fields struct {
		address         url.URL
		links           []url.URL
		lastFetchedDate time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   []url.URL
	}{
		0: {name: "ValidFields",
			fields: fields{address: fetchedURL(), links: nil, lastFetchedDate: sampleLastFetchedDate()},
			want:   nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			webPage := &WebPage{
				address:         tt.fields.address,
				links:           tt.fields.links,
				lastFetchedDate: tt.fields.lastFetchedDate,
			}
			if got := webPage.GetLinks(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLinks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWebPage_GetUrl(t *testing.T) {
	type fields struct {
		address         url.URL
		links           []url.URL
		lastFetchedDate time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   url.URL
	}{
		0: {name: "ValidFields",
			fields: fields{address: fetchedURL(), links: nil, lastFetchedDate: sampleLastFetchedDate()},
			want:   fetchedURL()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			webPage := &WebPage{
				address:         tt.fields.address,
				links:           tt.fields.links,
				lastFetchedDate: tt.fields.lastFetchedDate,
			}
			if got := webPage.GetUrl(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func fetchedURL() (url url.URL) {

	urlPtr, _ := url.Parse(fetchedUrlString)
	url = *urlPtr

	return url
}

func sampleLastFetchedDate() time.Time {

	sampleTime, _ := time.Parse(time.Now().Format(time.UnixDate), time.UnixDate)
	sampleTime = sampleTime.AddDate(1, 1, 1)

	return sampleTime
}

func sampleWebPage() *WebPage {

	webPage, _ := NewWebPage(fetchedURL(), nil, sampleLastFetchedDate())

	return webPage
}
