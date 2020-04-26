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
	List() (*dto.Users, error)
	GetByID(id string) (*dto.User, error)
	GetPasswordByEmail(email string) (*dto.User, error)
	Create(u *dto.User) error
	Update(u *dto.User) error
	Delete(id string) error
}

type UserRepoGorm struct {
	DB *gorm.DB
}

func (r UserRepoGorm) List() (*dto.Users, error) {
	rows, err := config.DB.Raw("SELECT id, name, email, age FROM users WHERE deleted_at IS NULL order by created_at DESC").Rows()
	if err != nil {
		return nil, internalError.MakeError(internalError.DatabaseError, err.Error())
	}
	defer rows.Close()
	result := dto.Users{}

	for rows.Next() {
		user := dto.User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Age)
		if err != nil {
			return nil, internalError.MakeError(internalError.ScanError, err.Error())
		}
		result.Users = append(result.Users, user)
	}
	return &result, nil
}

func (r UserRepoGorm) Create(u *dto.User) error {
	mUser := convertUser(u)
	op := r.DB.Create(&mUser)
	if err := op.Error; err != nil {
		switch err.(type) {
		case *pq.Error:
			if pqErr := err.(*pq.Error); pqErr.Code == "23505" {
				return internalError.MakeError(internalError.UniqueKeyError, "Email already exists")
			}
			return internalError.MakeError(internalError.DatabaseError, err.Error())
		default:
			return internalError.MakeError(internalError.DatabaseError, err.Error())
		}

	}
	return nil
}

func (r UserRepoGorm) Update(u *dto.User) error {
	_, iError := r.GetByID(u.ID)
	if iError != nil {
		return iError
	}
	op := r.DB.Model(u).Select("name, age").Updates(&u)
	if err := op.Error; err != nil {
		return internalError.MakeError(internalError.DatabaseError, err.Error())
	}
	return nil
}

func (r UserRepoGorm) GetByID(id string) (*dto.User, error) {
	mUser := new(model.User)
	op := r.DB.First(&mUser, "id = ?", id)
	if err := op.Scan(&mUser).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, internalError.MakeError(internalError.NotFoundError, "User not found")
		}
		return nil, internalError.MakeError(internalError.DatabaseError, err.Error())
	}
	return formatUser(mUser), nil
}

func (r UserRepoGorm) GetPasswordByEmail(email string) (*dto.User, error) {
	mUser := new(model.User)
	op := r.DB.First(&mUser, "email = ?", email)
	if err := op.Scan(&mUser).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, internalError.MakeError(internalError.AuthError, "Invalid credentials")
		}
		return nil, internalError.MakeError(internalError.DatabaseError, err.Error())
	}
	return &dto.User{
		ID:       mUser.ID.String(),
		Password: &mUser.Password,
	}, nil
}

func (r UserRepoGorm) Delete(id string) error {
	_, iError := r.GetByID(id)
	if iError != nil {
		return iError
	}
	mUser := new(model.User)
	mUser.ID = uuid.FromStringOrNil(id)
	op := r.DB.Delete(&mUser)
	if err := op.Error; err != nil {
		return internalError.MakeError(internalError.DatabaseError, err.Error())
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
