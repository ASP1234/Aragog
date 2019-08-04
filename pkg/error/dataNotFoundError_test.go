package error

import (
	"reflect"
	"testing"
)

const dataNotFoundErrorMsg = "Sample Error Message"

func TestDataNotFoundError_Error(t *testing.T) {
	type fields struct {
		message string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		0: {name: "NonEmptyMsg", fields: fields{message: dataNotFoundErrorMsg}, want: dataNotFoundErrorMsg}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := &DataNotFoundError{
				message: tt.fields.message,
			}
			if got := err.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDataNotFoundError(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
		want *DataNotFoundError
	}{
		0: {name: "NonEmptyMsg", args: args{message: dataNotFoundErrorMsg}, want: NewDataNotFoundError(dataNotFoundErrorMsg)}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDataNotFoundError(tt.args.message); !reflect.DeepEqual(got.message, tt.want.message) {
				t.Errorf("NewDataNotFoundError() = %v, want %v", got, tt.want)
			}
		})
	}
}
