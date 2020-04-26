package model

type Diary struct {
	Base
	AuthorID   string
	AuthorName string
	Title      string
	Content    string `gorm:"type:varchar(10000)"`
}
