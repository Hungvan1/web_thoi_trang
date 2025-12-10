package constant

type Response struct {
	Message string         `json:"message,omitempty"`
	Data    interface{}    `json:"data,omitempty"`
	Error   *ErrorResponse `json:"error,omitempty"`
}

type ErrorResponse struct {
	Message string `json:"message,omitempty"`
	Code    int    `json:"code,omitempty"`
}
