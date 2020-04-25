package repository

import (
	"github.com/golpo/dto"
	"github.com/golpo/model"
	"github.com/jinzhu/gorm"
)

type UserRepo interface {
	Create(u *dto.User) (*model.User, error)
	//Update(u *dto.User) (*model.User, serviceError)
	//Delete(u *dto.User) serviceError
}

type UserRepoGorm struct {
	DB *gorm.DB
}

func (r UserRepoGorm) Create(u *dto.User) (*model.User, error) {
	mUser := model.User{
		Email:    u.Email,
		Password: *u.Password,
		Name:     u.Name,
		Age:      u.Age,
	}
	op := r.DB.Create(&mUser)
	if err := op.Error; err != nil {
		return nil, err
	}
	return &mUser, nil
}
