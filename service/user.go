package service

import (
	"github.com/golpo/dto"
	"github.com/golpo/internalError"
	"github.com/golpo/repository"
	util "github.com/golpo/service/util"
)

type UserService interface {
	ListUsers() (*dto.Users, error)
	CreateUser(u *dto.User) error
	UpdateUser(u *dto.User) error
	DeleteUser(id string) error
}

type UserServiceImpl struct {
	UserRepo repository.UserRepo
}

func (s UserServiceImpl) ListUsers() (*dto.Users, error) {
	return s.UserRepo.List()
}

func (s UserServiceImpl) CreateUser(u *dto.User) error {
	pStr, err := util.HashPassword(*u.Password)
	if err != nil {
		return internalError.MakeError(internalError.HashError, err.Error())
	}
	u.Password = &pStr
	if err := s.UserRepo.Create(u); err != nil {
		return err
	}
	return nil
}

func (s UserServiceImpl) UpdateUser(u *dto.User) error {
	err := s.UserRepo.Update(u)
	if err != nil {
		return err
	}
	return nil
}

func (s UserServiceImpl) DeleteUser(id string) error {
	err := s.UserRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
