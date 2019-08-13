// Package error provides various custom errors for enhancing the error handling depending upon the use case.
package error

// DependencyError representing errors related to dependency failure.
type DependencyError struct {
	message string
}

func NewDependencyError(message string) *DependencyError {

	err := new(DependencyError)
	err.message = message

	return err
}

func (err *DependencyError) Error() string {
	return err.message
}
