package fetcher

import (
	"net/url"
	"testing"
)

func TestNewHttpFetcher(t *testing.T) {

	tests := []struct {
		name    string
		wantErr bool
	}{
		0: {name: "ValidURL", wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewHttpFetcher()
			if (err != nil) != tt.wantErr {
				t.Errorf("NewHttpFetcher() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestHttpFetcher_Fetch(t *testing.T) {

	exampleURL, _ := url.Parse("https://monzo.com")
	invalidURL, _ := url.Parse("/about/")

	type args struct {
		seedUrl url.URL
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		0: {name: "ValidURL", args: args{seedUrl: *exampleURL}, wantErr: false},
		1: {name: "EmptyURL", args: args{seedUrl: url.URL{}}, wantErr: true},
		2: {name: "InvalidURL", args: args{seedUrl: *invalidURL}, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpFetcher, _ := NewHttpFetcher()
			_, err := httpFetcher.Fetch(tt.args.seedUrl)

			if (err != nil) != tt.wantErr {
				t.Errorf("Fetch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
