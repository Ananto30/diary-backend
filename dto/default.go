package dto

type SuccessResponse struct {
	Success bool `json:"success"`
}

func NewSuccessResponse() *SuccessResponse {
	return &SuccessResponse{Success: true}
}
