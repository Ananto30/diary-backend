package service

import (
	"github.com/golpo/dto"
	"github.com/golpo/internalError"
	"github.com/golpo/repository"
	util "github.com/golpo/service/util"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type AuthService interface {
	Login(req *dto.LoginRequest) (*dto.LoginResponse, *internalError.IError)
}

type AuthServiceImpl struct {
	UserRepo repository.UserRepo
}

func (s AuthServiceImpl) Login(req *dto.LoginRequest) (*dto.LoginResponse, *internalError.IError) {
	u, ierr := s.UserRepo.GetPasswordByEmail(req.Email)
	if ierr != nil {
		return nil, ierr
	}

	err := bcrypt.CompareHashAndPassword([]byte(*u.Password), []byte(*req.Password))
	if err != nil {
		log.Println(err)
		return nil, internalError.Error(internalError.AuthError, "Invalid credentials")
	}

	tkn, err := util.GenerateToken(u.ID)
	if err != nil {
		return nil, internalError.Error(internalError.JwtError, err.Error())
	}

	return &dto.LoginResponse{AccessToken: tkn}, nil

}
