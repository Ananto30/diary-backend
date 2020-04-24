package service

import (
	"github.com/gofiber/fiber"
	"github.com/golpo/config"
	"github.com/golpo/dto"
	"github.com/golpo/util"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx, req *dto.LoginRequest) {
	pStr := util.HashPassword(c, *req.Password)
	req.Password = &pStr
	res := config.DB.Raw("SELECT id, password FROM users WHERE email=$1", req.Email).Row()
	u := dto.User{}
	if res.Scan(&u.ID, &u.Password) != nil {
		util.LogWithTrack(c, res.Scan().Error())
		c.Status(403).Send("Invalid credentials")
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(*u.Password), []byte(*req.Password))

	tkn, err := util.GenerateToken(c, u.ID)
	if err != nil {
		c.Status(503).Send("Service unavailable")
	}
	if err := c.JSON(dto.LoginResponse{AccessToken: tkn}); err != nil {
		c.Status(500).Send(err)
		return
	}

}
