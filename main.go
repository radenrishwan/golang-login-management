package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"login-management/database"
	"login-management/helper"
	"login-management/model/web"
	"login-management/repository"
	"login-management/service"
)

func main() {
	err := godotenv.Load()
	helper.PanicIfError(err)

	db := database.GetDatabase()
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db)

	//userService.Register(context.Background(), web.RegisterUserRequest{
	//	Username: "raden",
	//	Password: "inipassword",
	//	Email:    "muhamadrishwan87@gmail.com",
	//})

	response := userService.Login(context.Background(), web.LoginUserRequest{
		Username: "raden",
		Password: "inipassword",
	})

	result, err := json.Marshal(response)
	helper.PanicIfError(err)

	users := new(web.UserSessionResponse)
	err = json.Unmarshal(result, &users)
	helper.PanicIfError(err)

	fmt.Println(string(result))
}
