// Package error provides various custom errors for enhancing the error handling depending upon the use case.
package error

// DataNotFoundError representing errors related to requested data being not present.
type DataNotFoundError struct {
	message string
}

func NewDataNotFoundError(message string) *DataNotFoundError {

	err := new(DataNotFoundError)
	err.message = message

	return err
}

func (err *DataNotFoundError) Error() string {
	return err.message
}
