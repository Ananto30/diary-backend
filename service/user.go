package service

import (
	"github.com/golpo/dto"
	"github.com/golpo/internalError"
	"github.com/golpo/repository"
	util2 "github.com/golpo/service/util"
	_ "github.com/lib/pq"
)

type UserService interface {
	ListUsers() (*dto.Users, *internalError.IError)
	CreateUser(u *dto.User) *internalError.IError
	UpdateUser(u *dto.User) *internalError.IError
	DeleteUser(id string) *internalError.IError
}

type UserServiceImpl struct {
	UserRepo repository.UserRepo
}

func (s UserServiceImpl) ListUsers() (*dto.Users, *internalError.IError) {
	return s.UserRepo.List()
}

func (s UserServiceImpl) CreateUser(u *dto.User) *internalError.IError {
	pStr, err := util2.HashPassword(*u.Password)
	if err != nil {
		return internalError.Error(internalError.HashError, err.Error())
	}
	u.Password = &pStr
	if ierr := s.UserRepo.Create(u); ierr != nil {
		return ierr
	}
	return nil
}

func (s UserServiceImpl) UpdateUser(u *dto.User) *internalError.IError {
	ierr := s.UserRepo.Update(u)
	if ierr != nil {
		return ierr
	}
	return nil
}

func (s UserServiceImpl) DeleteUser(id string) *internalError.IError {
	ierr := s.UserRepo.Delete(id)
	if ierr != nil {
		return ierr
	}
	return nil
}
