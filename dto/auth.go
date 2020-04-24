package dto


// LoginRequest struct
type LoginRequest struct {
	Email    string  `json:"email"`
	Password *string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}
