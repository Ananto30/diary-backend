package model

import "github.com/jinzhu/gorm"

type Diary struct {
	gorm.Model
	AuthorID   string
	AuthorName string
	Title      string
	Content    string `gorm:"type:varchar(10000)"`
}
