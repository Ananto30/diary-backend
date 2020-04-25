package model

// User model
type User struct {
	Base
	Name     string
	Email    string `gorm:"type:varchar(100);unique_index"`
	Password string
	Age      uint8
}
