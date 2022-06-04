package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"log"
	"login-management/database"
	"login-management/handler"
	"login-management/helper"
	"login-management/repository"
	"login-management/router"
	"login-management/service"
)

func main() {
	err := godotenv.Load()
	helper.PanicIfError(err)

	db := database.GetDatabase()
	db.Exec("delete from sessions")
	db.Exec("delete from users")

	app := fiber.New(fiber.Config{
		AppName:      "login-management",
		ErrorHandler: handler.ErrorHandler,
	})

	app.Use(logger.New())
	app.Use(recover.New())

	// TODO: add dependency injection

	// Register Repository
	userRepository := repository.NewUserRepository()
	sessionRepository := repository.NewSessionRepository()

	// Register Service
	userService := service.NewUserService(userRepository, sessionRepository, db)

	// Register Handler
	userHandler := handler.NewUserHandler(userService)

	// Register routes
	router.NewUserRouter(app, userHandler)

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"message": "App Running !!!",
		})
	})

	log.Fatal(app.Listen(":3000"))
}
