package repository

import (
	"github.com/golpo/config"
	"github.com/golpo/dto"
	"github.com/golpo/internalError"
	"github.com/golpo/model"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

type UserRepo interface {
	List() (*dto.Users, *internalError.IError)
	GetByID(id string) (*dto.User, *internalError.IError)
	Create(u *dto.User) *internalError.IError
	Update(u *dto.User) *internalError.IError
	Delete(id string) *internalError.IError
}

type UserRepoGorm struct {
	DB *gorm.DB
}

func (r UserRepoGorm) List() (*dto.Users, *internalError.IError) {
	rows, err := config.DB.Raw("SELECT id, name, email, age FROM users WHERE deleted_at IS NULL order by id").Rows()
	if err != nil {
		return nil, internalError.Error(internalError.DatabaseError, err.Error())
	}
	defer rows.Close()
	result := dto.Users{}

	for rows.Next() {
		user := dto.User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Age)
		if err != nil {
			return nil, internalError.Error(internalError.ScanError, err.Error())
		}
		result.Users = append(result.Users, user)
	}
	return &result, nil
}

func (r UserRepoGorm) Create(u *dto.User) *internalError.IError {
	mUser := convertUser(u)
	op := r.DB.Create(&mUser)
	if err := op.Error; err != nil {
		switch err.(type) {
		case *pq.Error:
			if pqErr := err.(*pq.Error); pqErr.Code == "23505" {
				return internalError.Error(internalError.UniqueKeyError, "Email already exists")
			}
			return internalError.Error(internalError.DatabaseError, err.Error())
		default:
			return internalError.Error(internalError.DatabaseError, err.Error())
		}

	}
	return nil
}

func (r UserRepoGorm) Update(u *dto.User) *internalError.IError {
	_, iError := r.GetByID(u.ID)
	if iError != nil {
		return iError
	}
	op := r.DB.Model(u).Select("name, age").Updates(&u)
	if err := op.Error; err != nil {
		return internalError.Error(internalError.DatabaseError, err.Error())
	}
	return nil
}

func (r UserRepoGorm) GetByID(id string) (*dto.User, *internalError.IError) {
	mUser := new(model.User)
	op := r.DB.First(&mUser, "id = ?", id)
	if err := op.Scan(&mUser).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, internalError.Error(internalError.NotFoundError, "User not found")
		}
		return nil, internalError.Error(internalError.DatabaseError, err.Error())
	}
	return formatUser(mUser), nil
}

func (r UserRepoGorm) Delete(id string) *internalError.IError {
	_, iError := r.GetByID(id)
	if iError != nil {
		return iError
	}
	mUser := new(model.User)
	mUser.ID = uuid.FromStringOrNil(id)
	op := r.DB.Delete(&mUser)
	if err := op.Error; err != nil {
		return internalError.Error(internalError.DatabaseError, err.Error())
	}
	return nil
}

func formatUser(mUser *model.User) *dto.User {
	return &dto.User{
		ID:    mUser.ID.String(),
		Name:  mUser.Name,
		Email: mUser.Email,
		Age:   mUser.Age,
	}
}

func convertUser(u *dto.User) *model.User {
	mUser := model.User{
		Email:    u.Email,
		Password: *u.Password,
		Name:     u.Name,
		Age:      u.Age,
	}
	return &mUser
}
