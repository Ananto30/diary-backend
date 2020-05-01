package server

import (
	"github.com/gofiber/fiber"
	"github.com/golpo/config"
	"github.com/golpo/handler"
	"github.com/golpo/middleware"
	"github.com/golpo/repository"
)

func InitiateRoutes(app *fiber.App) {

	userRepo := repository.UserRepoGorm{DB: config.DB}
	userHandler := handler.UserHandler{UserRepo: &userRepo}

	userGroup := app.Group("/api/user")
	userGroup.Use(middleware.Auth())
	userGroup.Get("/", userHandler.UserList)
	userGroup.Post("/", userHandler.CreateUser)
	userGroup.Put("/", userHandler.UpdateUser)
	userGroup.Delete("/", userHandler.DeleteUser)
	userGroup.Get("/me", userHandler.GetUserByAuthToken)
	userGroup.Get("/:id", userHandler.GetUserById)


	authHandler := handler.AuthHandler{UserRepo: &userRepo}

	authGroup := app.Group("/api/auth")
	authGroup.Post("/login", authHandler.Login)

	diaryRepo := repository.DiaryRepoGorm{DB: config.DB}
	diaryHandler := handler.DiaryHandler{DiaryRepo: &diaryRepo}

	diaryGroup := app.Group("/api/diary")
	diaryGroup.Use(middleware.Auth())
	diaryGroup.Get("/", diaryHandler.DiaryList)
	diaryGroup.Post("/", diaryHandler.CreateDiary)
	diaryGroup.Get("/user/:id", diaryHandler.GetDiaryByUserId)
	diaryGroup.Get("/:id", diaryHandler.GetDiaryById)


}
