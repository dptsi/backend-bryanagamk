package main

import (
	"bryanagamk/go-rest-miwt/app"
	"bryanagamk/go-rest-miwt/controller"
	"bryanagamk/go-rest-miwt/exception"
	"bryanagamk/go-rest-miwt/helper"
	"bryanagamk/go-rest-miwt/repository"
	"bryanagamk/go-rest-miwt/service"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"github.com/spf13/viper"
	"net/http"
)

func main() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	helper.PanicIfError(err)

	fmt.Println("Running Server: ", viper.Get("SERVER_ADDRESS"))

	db := app.NewDB(viper.Get("DB_DRIVER"), viper.Get("DB_SOURCE"), viper.Get("DB_HOST"), viper.Get("DB_USER"), viper.Get("DB_PASSWORD"), viper.Get("DB_NAME"), viper.Get("DB_PORT"))
	validate := validator.New()
	cutiRepository := repository.NewCutiRepository()
	cutiService := service.NewCutiService(cutiRepository, db, validate)
	cutiController := controller.NewCutiController(cutiService)

	router := httprouter.New()

	router.GET("/api/riwayat_cuti", cutiController.FindAll)
	router.GET("/api/riwayat_cuti/:cutiId", cutiController.FindById)
	router.POST("/api/riwayat_cuti", cutiController.Create)
	router.PUT("/api/riwayat_cuti/:cutiId", cutiController.Update)
	router.DELETE("/api/riwayat_cuti/:cutiId", cutiController.Delete)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    fmt.Sprint(viper.Get("SERVER_ADDRESS")),
		Handler: router,
	}

	err = server.ListenAndServe()
	helper.PanicIfError(err)
}
