package producer

import (
	"Aragog/internal/pkg/entity"
	"net/url"
	"reflect"
	"testing"
)

const exampleChannelSize = 10

func TestLocalProducer_produce(t *testing.T) {

	exampleURL,_ := url.Parse("http://monzo.com")
	exampleMessage,_ := message.New(*exampleURL)

	type fields struct {
		messageQueue chan message.Message
	}
	type args struct {
		msg message.Message
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		0: {name: "ValidMessage",
			fields: fields{messageQueue: make(chan message.Message, exampleChannelSize)},
			args: args{msg: *exampleMessage},
			wantErr: false},

		1: {name: "EmptyMessage",
			fields: fields{messageQueue: make(chan message.Message, exampleChannelSize)},
			args: args{msg: message.Message{}},
			wantErr: true}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			producer := &LocalProducer{
				messageQueue: tt.fields.messageQueue,
			}
			if err := producer.produce(tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("produce() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewLocalProducer(t *testing.T) {

	exampleLocalProducer,_ := New(make(chan message.Message, exampleChannelSize))

	type args struct {
		messageQueue chan message.Message
	}
	tests := []struct {
		name         string
		args         args
		wantProducer *LocalProducer
		wantErr      bool
	}{
		0: {name: "ValidMessageQueue",
			args: args{messageQueue: exampleLocalProducer.messageQueue},
			wantProducer: exampleLocalProducer,
			wantErr: false},

		1: {name: "NilMessageQueue",
			args: args{messageQueue: nil},
			wantProducer: nil,
			wantErr: true}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotProducer, err := New(tt.args.messageQueue)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (gotProducer != nil) && (tt.wantProducer != nil) {
				if !reflect.DeepEqual(gotProducer.messageQueue, tt.wantProducer.messageQueue) {
					t.Errorf("New() gotProducer = %v, want %v", gotProducer, tt.wantProducer)
				}
			}
		})
	}
}