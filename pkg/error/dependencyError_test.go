package error

import (
	"reflect"
	"testing"
)

const dependencyErrorMsg = "Sample Error Message"

func TestDependencyError_Error(t *testing.T) {
	type fields struct {
		message string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		0: {name: "NonEmptyMsg", fields: fields{message: dependencyErrorMsg}, want: dependencyErrorMsg}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := &DependencyError{
				message: tt.fields.message,
			}
			if got := err.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDependencyError(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
		want *DependencyError
	}{
		0: {name: "NonEmptyMsg", args: args{message: dependencyErrorMsg}, want: NewDependencyError(dependencyErrorMsg)}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDependencyError(tt.args.message); !reflect.DeepEqual(got.message, tt.want.message) {
				t.Errorf("NewDependencyError() = %v, want %v", got, tt.want)
			}
		})
	}
}
