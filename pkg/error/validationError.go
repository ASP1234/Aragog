// Package error provides various custom errors for enhancing the error handling depending upon the use case.
package error

// ValidationError representing errors related to validation checks.
type ValidationError struct {
	message string
}

func NewValidationError(message string) *ValidationError {

	err := new(ValidationError)
	err.message = message

	return err
}

func (err *ValidationError) Error() string {
	return err.message
}
