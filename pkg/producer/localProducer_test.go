package producer

import (
	"github.com/ASP1234/Aragog/pkg/entity"
	"net/url"
	"reflect"
	"testing"
)

const exampleChannelSize = 10

func TestNewLocalProducer(t *testing.T) {

	exampleLocalProducer, _ := NewLocalProducer(make(chan entity.Message, exampleChannelSize))

	type args struct {
		messageQueue chan entity.Message
	}

	tests := []struct {
		name         string
		args         args
		wantProducer *LocalProducer
		wantErr      bool
	}{
		0: {name: "ValidMessageQueue",
			args:         args{messageQueue: exampleLocalProducer.messageQueue},
			wantProducer: exampleLocalProducer,
			wantErr:      false},

		1: {name: "NilMessageQueue",
			args:         args{messageQueue: nil},
			wantProducer: nil,
			wantErr:      true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotProducer, err := NewLocalProducer(tt.args.messageQueue)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (gotProducer != nil) && (tt.wantProducer != nil) {
				if !reflect.DeepEqual(gotProducer.messageQueue, tt.wantProducer.messageQueue) {
					t.Errorf("NewMessage() gotProducer = %v, want %v", gotProducer, tt.wantProducer)
				}
			}
		})
	}
}

func TestLocalProducer_produce(t *testing.T) {

	exampleURL, _ := url.Parse("http://monzo.com")
	exampleMessage, _ := entity.NewMessage(exampleURL)

	type fields struct {
		messageQueue chan entity.Message
	}

	type args struct {
		msg entity.Message
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		0: {name: "ValidMessage",
			fields:  fields{messageQueue: make(chan entity.Message, exampleChannelSize)},
			args:    args{msg: *exampleMessage},
			wantErr: false},

		1: {name: "EmptyMessage",
			fields:  fields{messageQueue: make(chan entity.Message, exampleChannelSize)},
			args:    args{msg: entity.Message{}},
			wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			producer := &LocalProducer{
				messageQueue: tt.fields.messageQueue,
			}
			if err := producer.Produce(tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("produce() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
