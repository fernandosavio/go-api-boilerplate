package validation

import (
	"example.com/bizday-api/internal/api/response"
)

type validationError struct {
	Code    string
	Message string
}

// Implement errors interface
func (be *validationError) Error() string {
	return be.Code + ": " + be.Message
}

func (e *validationError) Is(target error) bool {
	t, ok := target.(*validationError)
	return ok && e.Code == t.Code
}

func (e *validationError) FieldError(location response.ErrorLocation, context response.ErrorContext) *response.FieldError {
	return &response.FieldError{
		Code:     e.Code,
		Message:  e.Message,
		Location: location,
		Context:  context,
	}
}
