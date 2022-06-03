package service

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"login-management/helper"
	"login-management/model/domain"
	"login-management/model/web"
	"login-management/repository"
)

type UserService interface {
	Register(ctx context.Context, request web.RegisterUserRequest) web.UserSessionResponse
	Login(ctx context.Context, request web.LoginUserRequest) web.UserSessionResponse
	Update(ctx context.Context, request web.UpdateUserRequest) web.UserResponse
	Delete(ctx context.Context, request web.DeleteUserRequest) web.UserResponse
}

type userService struct {
	repository.UserRepository
	repository.SessionRepository
	*sql.DB
}

func NewUserService(userRepository repository.UserRepository, sessionRepository repository.SessionRepository, DB *sql.DB) UserService {
	return &userService{UserRepository: userRepository, SessionRepository: sessionRepository, DB: DB}
}

func (service *userService) Register(ctx context.Context, request web.RegisterUserRequest) web.UserSessionResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	password := helper.GeneratePassword(request.Password)

	user := domain.User{
		Id:        uuid.NewString(),
		Email:     request.Email,
		Username:  request.Username,
		Password:  password,
		Role:      domain.USER,
		CreatedAt: helper.GenerateMilisTimeNow(),
		UpdatedAt: helper.GenerateMilisTimeNow(),
	}

	service.UserRepository.Save(ctx, tx, user)

	// save session to database
	session := domain.Session{
		Id:        uuid.NewString(),
		UserId:    user.Id,
		CreatedAt: helper.GenerateMilisTimeNow(),
		ExpiredAt: helper.GenerateMilisTimeWeek(),
	}

	service.SessionRepository.Save(ctx, tx, session)

	// generate token
	token, err := helper.NewTokenString(session)

	return web.NewUserSessionResponse(user, token)
}

func (service *userService) Update(ctx context.Context, request web.UpdateUserRequest) web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	findUser, err := service.UserRepository.FindById(ctx, tx, domain.User{
		Id: request.Id,
	})
	helper.PanicIfError(err)

	err = bcrypt.CompareHashAndPassword([]byte(findUser.Password), []byte(request.Password))
	if err != nil {
		panic("user password is not valid")
	}

	password := helper.GeneratePassword(request.Password)
	user := domain.User{
		Id:        request.Id,
		Email:     request.Email,
		Username:  request.Username,
		Password:  password,
		Role:      findUser.Role,
		CreatedAt: findUser.CreatedAt,
		UpdatedAt: helper.GenerateMilisTimeNow(),
	}

	service.UserRepository.Update(ctx, tx, user)

	return web.NewUserResponse(user)
}

func (service *userService) Delete(ctx context.Context, request web.DeleteUserRequest) web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	findUser, err := service.UserRepository.FindById(ctx, tx, domain.User{
		Id: request.Id,
	})
	helper.PanicIfError(err)

	err = bcrypt.CompareHashAndPassword([]byte(findUser.Password), []byte(request.Password))
	if err != nil {
		panic("user password is not valid")
	}

	user := domain.User{
		Id: request.Id,
	}

	session := domain.Session{
		UserId: user.Id,
	}

	// delete from repository
	service.UserRepository.Delete(ctx, tx, user)
	service.SessionRepository.Delete(ctx, tx, session)

	return web.NewUserResponse(findUser)
}

func (service *userService) Login(ctx context.Context, request web.LoginUserRequest) web.UserSessionResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	findUser, err := service.UserRepository.FindByUsername(ctx, tx, domain.User{Username: request.Username})
	helper.PanicIfError(err)

	err = bcrypt.CompareHashAndPassword([]byte(findUser.Password), []byte(request.Password))
	if err != nil {
		panic("user password is not valid")
	}

	// check if session exist
	session, err := service.SessionRepository.FindById(ctx, tx, domain.Session{
		UserId: findUser.Id,
	})

	token := ""
	if err != nil {
		// if session not exist, create new session
		session := domain.Session{
			Id:        uuid.NewString(),
			UserId:    findUser.Id,
			CreatedAt: helper.GenerateMilisTimeNow(),
			ExpiredAt: helper.GenerateMilisTimeWeek(),
		}

		service.SessionRepository.Save(ctx, tx, session)

		token, err = helper.NewTokenString(session)
		helper.PanicIfError(err)
	} else {
		token, err = helper.NewTokenString(session)
		helper.PanicIfError(err)
	}

	return web.NewUserSessionResponse(findUser, token)
}
