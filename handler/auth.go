package handler

import (
	"github.com/gofiber/fiber"
	"github.com/golpo/dto"
	"github.com/golpo/internalError"
	"github.com/golpo/repository"
	"github.com/golpo/util"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	UserRepo repository.UserRepo
}

func (h AuthHandler) Login(c *fiber.Ctx) {
	req := new(dto.LoginRequest)
	if err := c.BodyParser(req); err != nil {
		c.Status(400).Send(err)
		return
	}

	u, ierr := h.UserRepo.GetPasswordByEmail(req.Email)
	if ierr != nil {
		errorHandler(c, ierr)
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(*u.Password), []byte(*req.Password))
	if err != nil {
		errorHandler(c, internalError.MakeError(internalError.AuthError, "Invalid credentials"))
		return
	}

	tkn, err := util.GenerateToken(c, u.ID)
	if err != nil {
		errorHandler(c, internalError.MakeError(internalError.JwtError, "Invalid credentials"))
		return
	}
	c.JSON(dto.LoginResponse{AccessToken: tkn})
}
