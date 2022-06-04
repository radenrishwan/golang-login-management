package handler

import (
	"github.com/gofiber/fiber/v2"
	"login-management/common"
	"login-management/exception"
	"net/http"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	_, ok := err.(exception.ValidateException)
	if ok {
		return ctx.Status(http.StatusBadRequest).JSON(common.Response[string]{
			Code:    http.StatusBadRequest,
			Message: "error",
			Data:    err.Error(),
		})
	}

	_, ok = err.(exception.UserException)
	if ok {
		return ctx.Status(http.StatusBadRequest).JSON(common.Response[string]{
			Code:    http.StatusBadRequest,
			Message: "error",
			Data:    err.Error(),
		})
	}

	return ctx.Status(http.StatusInternalServerError).JSON(common.Response[string]{
		Code:    http.StatusInternalServerError,
		Message: "error",
		Data:    err.Error(),
	})
}
