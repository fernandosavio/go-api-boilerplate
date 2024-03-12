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

type ErrorLocation string

const (
	LocPath        ErrorLocation = "path"
	LocQuerystring ErrorLocation = "querystring"
	LocHeader      ErrorLocation = "header"
	LocBody        ErrorLocation = "body"
)

type ValidationErrorResponse struct {
	Code      string                  `json:"code"`
	RequestID string                  `json:"request_id"`
	Errors    map[string][]FieldError `json:"errors"`
}

func (ve *ValidationErrorResponse) HasErrors() bool {
	return len(ve.Errors) > 0
}

func (ve *ValidationErrorResponse) Add(field string, error *FieldError) {
	if ve.Errors == nil {
		ve.Errors = make(map[string][]FieldError)
	}

	ve.Errors[field] = append(ve.Errors[field], *error)
}

type ErrorContext map[string]any

type FieldError struct {
	Code     string        `json:"code"`
	Location ErrorLocation `json:"location"`
	Message  string        `json:"message"`
	Context  ErrorContext  `json:"context,omitempty"`
}
