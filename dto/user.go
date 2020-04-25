package dto

// User struct
type User struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Password *string `json:"password"`
	Age      uint8   `json:"age"`
}

// Users struct
type Users struct {
	Users []User `json:"users"`
}
