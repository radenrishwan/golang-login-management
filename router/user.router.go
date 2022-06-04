package router

import (
	"github.com/gofiber/fiber/v2"
	"login-management/handler"
)

func NewUserRouter(app *fiber.App, userHandler handler.UserHandler) {
	app.Post("/api/v1/user/register", userHandler.Register)
	app.Post("/api/v1/user/login", userHandler.Login)
	app.Delete("/api/v1/user/delete", userHandler.Delete)
	app.Put("/api/v1/user/update", userHandler.Update)
}
