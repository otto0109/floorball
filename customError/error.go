package responseError

import "fmt"

// New returns an error that formats as the given text.
func New(code int, text string) *ResponseError {
	return &ResponseError{code, text}
}

func (err *ResponseError) Error() string {
	return fmt.Sprintf("Vicci api returned status: %v with code %v", err.statusCode, err.errorText)
}

type ResponseError struct {
	statusCode int
	errorText  string
}
