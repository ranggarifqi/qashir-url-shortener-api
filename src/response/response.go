package response

// ErrorResponse object
type ErrorResponse struct {
	StatusCode int
	Message    string
}

// SuccessResponse ...
type SuccessResponse struct {
	StatusCode int
	Message    string
	Data       interface{}
}
