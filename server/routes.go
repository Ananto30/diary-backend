package server

import (
	"github.com/gofiber/fiber"
	"github.com/golpo/config"
	"github.com/golpo/handler"
	"github.com/golpo/middleware"
	"github.com/golpo/repository"
	"github.com/golpo/service"
)

func InitiateRoutes(app *fiber.App) {

	userRepo := repository.UserRepoGorm{DB: config.DB}
	userService := service.UserServiceImpl{UserRepo: &userRepo}
	userHandler := handler.UserHandler{UserService: &userService}

	userGroup := app.Group("/api/user")
	userGroup.Use(middleware.Auth())
	userGroup.Get("/", userHandler.UserList)
	userGroup.Post("/", userHandler.CreateUser)
	userGroup.Put("/", userHandler.UpdateUser)
	userGroup.Delete("/", userHandler.DeleteUser)

	authService := service.AuthServiceImpl{UserRepo: &userRepo}
	authHandler := handler.AuthHandler{AuthService: &authService}

	authGroup := app.Group("/api/auth")
	authGroup.Post("/login", authHandler.Login)

	diaryRepo := repository.DiaryRepoGorm{DB: config.DB}
	diaryService := service.DiaryServiceImpl{DiaryRepo: &diaryRepo}
	diaryHandler := handler.DiaryHandler{DiaryService: &diaryService}

	diaryGroup := app.Group("/api/diary")
	diaryGroup.Use(middleware.Auth())
	diaryGroup.Get("/", diaryHandler.DiaryList)
	diaryGroup.Post("/", diaryHandler.CreateDiary)
}
