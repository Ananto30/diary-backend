package dto

import "time"

// Diary struct
type Diary struct {
	ID         string    `json:"id"`
	AuthorID   string    `json:"author_id"`
	AuthorName string    `json:"author_name"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"created_at"`
}

// Diaries struct
type Diaries struct {
	Diaries []Diary `json:"diaries"`
}
