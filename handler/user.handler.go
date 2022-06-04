package handler

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"login-management/common"
	"login-management/helper"
	"login-management/model/web"
	"login-management/service"
	"net/http"
)

type UserHandler interface {
	Register(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
}

type userHandler struct {
	service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return &userHandler{UserService: userService}
}

func (handler *userHandler) Register(ctx *fiber.Ctx) error {
	var request web.RegisterUserRequest

	err := ctx.BodyParser(&request)
	helper.PanicIfError(err)

	result := handler.UserService.Register(context.Background(), request)

	return ctx.Status(http.StatusCreated).JSON(common.Response[web.UserSessionResponse]{
		Code:    http.StatusCreated,
		Message: "User registered successfully",
		Data:    result,
	})
}

func (handler *userHandler) Login(ctx *fiber.Ctx) error {
	var request web.LoginUserRequest

	err := ctx.BodyParser(&request)
	helper.PanicIfError(err)

	result := handler.UserService.Login(context.Background(), request)

	return ctx.Status(http.StatusOK).JSON(common.Response[web.UserSessionResponse]{
		Code:    http.StatusOK,
		Message: "User logged in successfully",
		Data:    result,
	})
}

func (handler *userHandler) Delete(ctx *fiber.Ctx) error {
	var request web.DeleteUserRequest

	err := ctx.BodyParser(&request)
	helper.PanicIfError(err)

	result := handler.UserService.Delete(context.Background(), request)

	return ctx.Status(http.StatusOK).JSON(common.Response[web.UserResponse]{
		Code:    http.StatusOK,
		Message: "User delete successfully",
		Data:    result,
	})
}

func (handler *userHandler) Update(ctx *fiber.Ctx) error {
	var request web.UpdateUserRequest

	err := ctx.BodyParser(&request)
	helper.PanicIfError(err)

	result := handler.UserService.Update(context.Background(), request)

	return ctx.Status(http.StatusOK).JSON(common.Response[web.UserResponse]{
		Code:    http.StatusOK,
		Message: "User update successfully",
		Data:    result,
	})
}
