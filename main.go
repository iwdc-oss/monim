package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/iwdc-oss/monim/server/app"
	"github.com/iwdc-oss/monim/server/controller"
	"github.com/iwdc-oss/monim/server/helper"
	"github.com/iwdc-oss/monim/server/repository"
	"github.com/iwdc-oss/monim/server/service"
)

func main() {

	DB := app.NewDB()
	validate := validator.New()

	categoryRepository := repository.NewUserRepository()
	categoryService := service.NewUserService(categoryRepository, DB, validate)
	categoryController := controller.NewUserController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
