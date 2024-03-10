package validations

type ValidationError struct {
	code    string
	message string
}

func (be *ValidationError) Error() string {
	return be.code + ": " + be.message
}

func (e *ValidationError) Is(target error) bool {
	t, ok := target.(*ValidationError)
	return ok && e.code == t.code
}

var InvalidDateError = &ValidationError{
	code:    "invalid-date",
	message: "invalid date format, should be YYYY-MM-DD.",
}

var InvalidServiceError = &ValidationError{
	code:    "invalid-holiday-service",
	message: "service must not be nil",
}
