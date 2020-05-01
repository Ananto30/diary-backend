package repository

import (
	"github.com/golpo/config"
	"github.com/golpo/dto"
	"github.com/golpo/internalError"
	"github.com/golpo/model"
	"github.com/jinzhu/gorm"
)

type DiaryRepo interface {
	GetByID(id string) (*dto.Diary, error)
	GetByUserID(uID string) (*dto.Diaries, error)
	List() (*dto.Diaries, error)
	Create(d *dto.Diary) error
}

type DiaryRepoGorm struct {
	DB *gorm.DB
}

func (r DiaryRepoGorm) List() (*dto.Diaries, error) {
	rows, err := config.DB.Raw("SELECT id, title, author_id, content, created_at FROM diaries order by created_at DESC").Rows()
	if err != nil {
		return nil, internalError.MakeError(internalError.DatabaseError, err.Error())
	}
	defer rows.Close()
	result := dto.Diaries{}

	for rows.Next() {
		diary := dto.Diary{}
		err := rows.Scan(&diary.ID, &diary.Title, &diary.AuthorID, &diary.Content, &diary.CreatedAt)
		if err != nil {
			return nil, internalError.MakeError(internalError.ScanError, err.Error())
		}
		result.Diaries = append(result.Diaries, diary)
	}
	return &result, nil
}

func (r DiaryRepoGorm) Create(d *dto.Diary) error {
	mDiary := model.Diary{
		AuthorID:   d.AuthorID,
		AuthorName: d.AuthorName,
		Title:      d.Title,
		Content:    d.Content,
	}
	op := r.DB.Create(&mDiary)
	if err := op.Error; err != nil {
		return internalError.MakeError(internalError.DatabaseError, err.Error())
	}
	return nil
}

func (r DiaryRepoGorm) GetByID(id string) (*dto.Diary, error) {
	mD := new(model.Diary)
	op := r.DB.First(&mD, "id = ?", id)
	if err := op.Scan(&mD).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, internalError.MakeError(internalError.NotFoundError, "Diary not found")
		}
		return nil, internalError.MakeError(internalError.DatabaseError, err.Error())
	}
	return &dto.Diary{
		ID:        mD.ID.String(),
		AuthorID:  mD.AuthorID,
		Title:     mD.Title,
		Content:   mD.Content,
		CreatedAt: mD.CreatedAt,
	}, nil
}

func (r DiaryRepoGorm) GetByUserID(uID string) (*dto.Diaries, error) {
	rows, err := config.DB.Raw("SELECT id, title, author_id, content, created_at FROM diaries WHERE author_id=$1 order by created_at DESC", uID).Rows()
	if err != nil {
		return nil, internalError.MakeError(internalError.DatabaseError, err.Error())
	}
	defer rows.Close()
	result := dto.Diaries{}

	for rows.Next() {
		diary := dto.Diary{}
		err := rows.Scan(&diary.ID, &diary.Title, &diary.AuthorID, &diary.Content, &diary.CreatedAt)
		if err != nil {
			return nil, internalError.MakeError(internalError.ScanError, err.Error())
		}
		result.Diaries = append(result.Diaries, diary)
	}
	return &result, nil
}
