package message

import (
	"net/url"
	"reflect"
	"testing"
)

const exampleUrlString = "https://monzo.com/"

func TestMessage_GetLink(t *testing.T) {

	type fields struct {
		link url.URL
	}
	tests := []struct {
		name   string
		fields fields
		want   url.URL
	}{
		0: {name: "ValidURL", fields: fields{link: exampleURL()}, want: exampleURL()}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := &Message{
				link: tt.fields.link,
			}
			if got := msg.GetLink(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLink() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewMessage(t *testing.T) {
	type args struct {
		url url.URL
	}
	tests := []struct {
		name    string
		args    args
		wantMsg *Message
		wantErr bool
	}{
		0: {name: "ValidURL", args: args{url: exampleURL()}, wantMsg: &Message{link: exampleURL()}, wantErr: false},
		1: {name: "EmptyURL", args: args{url: url.URL{}}, wantMsg: nil, wantErr: true}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMsg, err := New(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotMsg, tt.wantMsg) {
				t.Errorf("New() gotMsg = %v, want %v", gotMsg, tt.wantMsg)
			}
		})
	}
}

func exampleURL() (url url.URL) {

	urlPtr, _ := url.Parse(exampleUrlString)
	url = *urlPtr

	return url
}