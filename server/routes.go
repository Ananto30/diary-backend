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
	//userService := service.UserServiceImpl{UserRepo: &userRepo}
	//userHandler := handler.UserHandler{UserService: &userService}
	//
	//userGroup := app.Group("/api/user")
	//userGroup.Use(middleware.Auth())
	//userGroup.Get("/", userHandler.UserList)
	//userGroup.Post("/", userHandler.CreateUser)
	//userGroup.Put("/", userHandler.UpdateUser)
	//userGroup.Delete("/", userHandler.DeleteUser)
	//

	authHandler := handler.AuthHandler{UserRepo: &userRepo}

	authGroup := app.Group("/api/auth")
	authGroup.Post("/login", authHandler.Login)
	//
	//diaryRepo := repository.DiaryRepoGorm{DB: config.DB}
	//diaryService := service.DiaryServiceImpl{DiaryRepo: &diaryRepo}
	//diaryHandler := handler.DiaryHandler{DiaryService: &diaryService}
	//
	//diaryGroup := app.Group("/api/diary")
	//diaryGroup.Use(middleware.Auth())
	//diaryGroup.Get("/", diaryHandler.DiaryList)
	//diaryGroup.Post("/", diaryHandler.CreateDiary)

	userGroup := app.Group("/api/user")
	userGroup.Use(middleware.Auth())
	userGroup.Get("/", handler.UserList)
	userGroup.Post("/", handler.CreateUser)
	userGroup.Put("/", handler.UpdateUser)
	userGroup.Delete("/", handler.DeleteUser)

	diaryGroup := app.Group("/api/diary")
	diaryGroup.Use(middleware.Auth())
	diaryGroup.Get("/", handler.DiaryList)
	diaryGroup.Post("/", handler.CreateDiary)
}
