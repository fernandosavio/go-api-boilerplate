package validation

var RequiredError = &validationError{
	Code:    "required",
	Message: "mandatory field value not informed",
}

var InvalidDateError = &validationError{
	Code:    "invalid-date",
	Message: "invalid date, should be a valid date in the format YYYY-MM-DD",
}

var InvalidHolidayServiceError = &validationError{
	Code:    "invalid-holiday-service",
	Message: "service must not be nil",
}

var MinValueError = &validationError{
	Code:    "min-value",
	Message: "field value is less then the minimal accepted value",
}

var MaxValueError = &validationError{
	Code:    "max-value",
	Message: "field value is more then the maximum accepted value",
}

var MinLengthError = &validationError{
	Code:    "min-length",
	Message: "length is less then the minimal accepted value",
}

var MaxLengthError = &validationError{
	Code:    "max-length",
	Message: "length is more then the maximum accepted value",
}
