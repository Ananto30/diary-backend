package repository

import (
	"github.com/golpo/config"
	"github.com/golpo/dto"
	"github.com/golpo/internalError"
	"github.com/golpo/model"
	"github.com/jinzhu/gorm"
)

type DiaryRepo interface {
	List() (*dto.Diaries, *internalError.IError)
	Create(d *dto.Diary) *internalError.IError
}

type DiaryRepoGorm struct {
	DB *gorm.DB
}

func (r DiaryRepoGorm) List() (*dto.Diaries, *internalError.IError) {
	rows, err := config.DB.Raw("SELECT id, title, author_id, content, created_at FROM diaries order by created_at").Rows()
	if err != nil {
		return nil, internalError.Error(internalError.DatabaseError, err.Error())
	}
	defer rows.Close()
	result := dto.Diaries{}

	for rows.Next() {
		diary := dto.Diary{}
		err := rows.Scan(&diary.ID, &diary.Title, &diary.AuthorID, &diary.Content, &diary.CreatedAt)
		if err != nil {
			return nil, internalError.Error(internalError.ScanError, err.Error())
		}
		result.Diaries = append(result.Diaries, diary)
	}
	return &result, nil
}

func (r DiaryRepoGorm) Create(d *dto.Diary) *internalError.IError {
	mDiary := model.Diary{
		AuthorID:   d.AuthorID,
		AuthorName: d.AuthorName,
		Title:      d.Title,
		Content:    d.Content,
	}
	op := r.DB.Create(&mDiary)
	if err := op.Error; err != nil {
		return internalError.Error(internalError.DatabaseError, err.Error())
	}
	return nil
}
