package service

import (
	"github.com/gofiber/fiber"
	"github.com/golpo/config"
	"github.com/golpo/dto"
	"github.com/golpo/repository"
	util2 "github.com/golpo/service/util"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"log"
)

type UserService interface {
	ListUsers() dto.Users
	CreateUser(u *dto.User) *dto.User
	UpdateUser(u *dto.User) *dto.User
	DeleteUser(id string)
}

type UserServiceImpl struct {
	UserRepo repository.UserRepo
}

func (s UserServiceImpl) GetUsers() (*dto.Users, *SError) {
	rows, err := config.DB.Raw("SELECT id, name, email, age FROM users order by id").Rows()
	if err != nil {
		return nil, Error(DatabaseError, err.Error())
	}
	defer rows.Close()
	result := dto.Users{}

	for rows.Next() {
		user := dto.User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Age)
		if err != nil {
			return nil, Error(ScanError, err.Error())
		}
		result.Users = append(result.Users, user)
	}

	return &result, nil
}

func (s UserServiceImpl) CreateUser(u *dto.User) (*dto.User, *SError) {
	pStr, err := util2.HashPassword(*u.Password)
	if err != nil {
		return nil, Error(HashError, err.Error())
	}
	u.Password = &pStr
	//res := config.DB.Exec("INSERT INTO users (name, email, password, age)VALUES ($1, $2, $3, $4)", u.Name, u.Email, pStr, u.Age)
	m, err := s.UserRepo.Create(u)

	pqErr := err.(*pq.Error)
	log.Println(pqErr.Code)
	if err != nil {
		return nil, Error(DatabaseError, err.Error())
	}
	log.Println(m)
	return u, nil
}

func UpdateUser(c *fiber.Ctx, u *dto.User) {
	//res := config.DB.Exec("UPDATE users SET name=$1,age=$2 WHERE id=$3", u.Name, u.Age, u.ID)
	//if err := res.Error; err != nil {
	//	util.LogWithTrack(c, err.Error())
	//	c.Status(500).JSON(dto.ServerError(c))
	//	return
	//}
	//if err := c.Status(201).JSON(dto.StatusResponse{Status: "Updated"}); err != nil {
	//	util.LogWithTrack(c, err.Error())
	//	c.Status(500).JSON(dto.ServerError(c))
	//	return
	//}
}

func DeleteUser(c *fiber.Ctx, u *dto.User) {
	//res := config.DB.Exec("DELETE FROM users WHERE id = $1", u.ID)
	//mU := &model.User{}
	//mU.ID = u.ID
	//res := config.DB.Delete()
	//if err := res.Error; err != nil {
	//	util.LogWithTrack(c, err.Error())
	//	c.Status(500).JSON(dto.ServerError(c))
	//	return
	//}
	//if err := c.JSON(dto.StatusResponse{Status: "Deleted"}); err != nil {
	//	util.LogWithTrack(c, err.Error())
	//	c.Status(500).JSON(dto.ServerError(c))
	//	return
	//}
}
