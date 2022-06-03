package service

import (
	"context"
	"database/sql"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"login-management/database"
	"login-management/helper"
	"login-management/model/domain"
	"login-management/model/web"
	"login-management/repository"
	"login-management/service"
	"testing"
)

func setupDb() *sql.DB {
	err := godotenv.Load("../../.env")
	helper.PanicIfError(err)

	return database.GetDatabase()
}

func truncateDatabase(db *sql.DB) {
	_, err := db.Exec("DELETE FROM sessions")
	helper.PanicIfError(err)

	_, err = db.Exec("DELETE FROM users")
	helper.PanicIfError(err)
}

func setupRepository() (repository.UserRepository, repository.SessionRepository) {
	userRepository := repository.NewUserRepository()
	sessionRepository := repository.NewSessionRepository()

	return userRepository, sessionRepository
}

func setupService() service.UserService {
	truncateDatabase(setupDb())

	db := database.GetDatabase()
	userRepository, sessionRepository := setupRepository()

	return service.NewUserService(userRepository, sessionRepository, db)
}

func TestRegisterSuccess(t *testing.T) {
	db := setupDb()
	userService := setupService()
	userRepository, sessionRepository := setupRepository()

	tx, err := db.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	register := userService.Register(context.Background(), web.RegisterUserRequest{
		Username: "raden",
		Password: "inipassword",
		Email:    "testemail@gmail.com",
	})

	pwd, err := bcrypt.GenerateFromPassword([]byte(register.Password), bcrypt.DefaultCost)

	assert.Equal(t, register.Username, "raden")
	assert.Nil(t, bcrypt.CompareHashAndPassword(pwd, []byte(register.Password)))
	assert.Equal(t, register.Email, "testemail@gmail.com")

	assert.NotNil(t, register.Id)

	result, err := userRepository.FindById(context.Background(), tx, domain.User{
		Id: register.Id,
	})

	assert.Nil(t, err)
	assert.Equal(t, result.Username, "raden")

	_, err = sessionRepository.FindById(context.Background(), tx, domain.Session{UserId: register.Id})

	assert.Nil(t, err)
}
