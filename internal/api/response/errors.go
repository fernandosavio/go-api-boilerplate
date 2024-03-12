package response

func Error404(correlationID string) *ResponseError {
	return &ResponseError{
		Code:      "not-found",
		Message:   "resource not found",
		RequestID: correlationID,
	}
}

func ValidationError(correlationId string, errors map[string][]FieldError) *ValidationErrorResponse {
	return &ValidationErrorResponse{
		Code:      "validation-error",
		RequestID: correlationId,
		Errors:    errors,
	}
}
