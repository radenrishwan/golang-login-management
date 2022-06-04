package validation

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"login-management/exception"
	"login-management/model/web"
)

const (
	length string = "lenght must be between %d and %d"
)

func UserRegisterValidation(request web.RegisterUserRequest) {
	err := validation.ValidateStruct(
		&request,
		validation.Field(&request.Username,
			validation.Required,
			validation.Length(3, 18).Error(fmt.Sprintf(length, 3, 18)),
		),
		validation.Field(&request.Password,
			validation.Required,
			validation.Length(3, 18).Error(fmt.Sprintf(length, 3, 18)),
		),
		validation.Field(&request.Email,
			validation.Required,
			validation.Length(3, 28).Error(fmt.Sprintf(length, 3, 18)),
			is.Email.Error("email must be valid"),
		),
	)

	if err != nil {
		panic(exception.NewValidateException(err.Error()))
	}
}

func UserLoginValidation(request web.LoginUserRequest) {
	err := validation.ValidateStruct(
		&request,
		validation.Field(&request.Username,
			validation.Required,
			validation.Length(3, 18).Error(fmt.Sprintf(length, 3, 18)),
		),
		validation.Field(&request.Password,
			validation.Required,
			validation.Length(3, 18).Error(fmt.Sprintf(length, 3, 18)),
		),
	)

	if err != nil {
		panic(exception.NewValidateException(err.Error()))
	}
}

func UserDeleteValidation(request web.DeleteUserRequest) {
	err := validation.ValidateStruct(
		&request,
		validation.Field(&request.Id,
			validation.Required,
		),
		validation.Field(&request.Password,
			validation.Required,
			validation.Length(3, 18).Error(fmt.Sprintf(length, 3, 18)),
		),
	)

	if err != nil {
		panic(exception.NewValidateException(err.Error()))
	}
}

func UserUpdateValidation(request web.UpdateUserRequest) {
	err := validation.ValidateStruct(
		&request,
		validation.Field(&request.Id,
			validation.Required,
		),
		validation.Field(&request.Username,
			validation.Required,
			validation.Length(3, 18).Error(fmt.Sprintf(length, 3, 18)),
		),
		validation.Field(&request.Password,
			validation.Required,
			validation.Length(3, 18).Error(fmt.Sprintf(length, 3, 18)),
		),
		validation.Field(&request.Email,
			validation.Required,
			validation.Length(3, 28).Error(fmt.Sprintf(length, 3, 18)),
			is.Email.Error("email must be valid"),
		),
	)

	if err != nil {
		panic(exception.NewValidateException(err.Error()))
	}
}
