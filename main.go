package main

import (
	"github.com/joho/godotenv"
	"login-management/helper"
)

func main() {
	err := godotenv.Load()
	helper.PanicIfError(err)
}
