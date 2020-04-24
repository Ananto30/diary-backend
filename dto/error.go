package dto

type ErrorResponse struct {
	RequestID string `json:"request_id"`
	ErrorCode int16 `json:"error_code"`

}
