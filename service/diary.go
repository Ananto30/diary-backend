package service

import (
	"github.com/golpo/dto"
	"github.com/golpo/internalError"
	"github.com/golpo/repository"
)

type DiaryService interface {
	ListDiaries() (*dto.Diaries, *internalError.IError)
	CreateDiary(d *dto.Diary) *internalError.IError
}

type DiaryServiceImpl struct {
	DiaryRepo repository.DiaryRepo
}

func (s DiaryServiceImpl) ListDiaries() (*dto.Diaries, *internalError.IError) {
	return s.DiaryRepo.List()
}

func (s DiaryServiceImpl) CreateDiary(d *dto.Diary) *internalError.IError {
	if ierr := s.DiaryRepo.Create(d); ierr != nil {
		return ierr
	}
	return nil
}
