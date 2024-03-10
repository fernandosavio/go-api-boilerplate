package response

type ResponseError struct {
	Code      string `json:"code"`
	Message   string `json:"message"`
	RequestID string `json:"request_id"`
}

func (re *ResponseError) Error() string {
	return re.Message
}

func (re *ResponseError) Is(target error) bool {
	t, ok := target.(*ResponseError)
	return ok && re.Code == t.Code
}

func Error404(correlationID string) *ResponseError {
	return &ResponseError{
		Code:      "not-found",
		Message:   "resource not found",
		RequestID: correlationID,
	}
}
