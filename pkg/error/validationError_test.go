package error

import (
	"reflect"
	"testing"
)

const validationErrorMsg = "Sample Error Message"

func TestValidationError_Error(t *testing.T) {
	type fields struct {
		message string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		0: {name: "NonEmptyMsg", fields: fields{message: validationErrorMsg}, want: validationErrorMsg}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := &ValidationError{
				message: tt.fields.message,
			}
			if got := err.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewValidationError(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
		want *ValidationError
	}{
		0: {name: "NonEmptyMsg", args: args{message: validationErrorMsg}, want: NewValidationError(validationErrorMsg)}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewValidationError(tt.args.message); !reflect.DeepEqual(got.message, tt.want.message) {
				t.Errorf("NewValidationError() = %v, want %v", got, tt.want)
			}
		})
	}
}
