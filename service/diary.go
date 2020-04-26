package service

import (
	"github.com/golpo/dto"
	"github.com/golpo/repository"
)

type DiaryService interface {
	ListDiaries() (*dto.Diaries, error)
	CreateDiary(d *dto.Diary) error
}

type DiaryServiceImpl struct {
	DiaryRepo repository.DiaryRepo
}

func (s DiaryServiceImpl) ListDiaries() (*dto.Diaries, error) {
	return s.DiaryRepo.List()
}

func (s DiaryServiceImpl) CreateDiary(d *dto.Diary) error {
	if err := s.DiaryRepo.Create(d); err != nil {
		return err
	}
	return nil
}
