package shared

type ErrorResponse struct {
	CustomErrorCode string `json:"customErrorCode,omitempty"`
	Message         string `json:"message,omitempty"`
}
