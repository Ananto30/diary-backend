package model

import "github.com/jinzhu/gorm"

// User model
type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"type:varchar(100);unique_index"`
	Password string
	Age      string
}
