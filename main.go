package main

import (
	"fmt"
	"github.com/IndarMuis/pt-xyz.git/src/common/config"
	"github.com/IndarMuis/pt-xyz.git/src/common/exception"
	"github.com/IndarMuis/pt-xyz.git/src/common/utils"
	route "github.com/IndarMuis/pt-xyz.git/src/controller"
	controller "github.com/IndarMuis/pt-xyz.git/src/controller/impl"
	repository "github.com/IndarMuis/pt-xyz.git/src/repository/impl"
	service "github.com/IndarMuis/pt-xyz.git/src/service/impl"
	validator2 "github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func main() {
	//init database
	DB := config.NewDatabase()

	validator := validator2.New()

	// init customer
	customerRepository := repository.NewCustomerRepository()
	customerService := service.NewCustomerService(customerRepository, DB, validator)
	customerController := controller.NewCustomerController(customerService)

	// init customer tenor
	customerTenorRepository := repository.NewCustomerTenorRepository()
	customerTenorService := service.NewCustomerTenorService(customerTenorRepository, DB, validator)
	customerTenorController := controller.NewCustomerTenorController(customerTenorService)

	// init customer transaction
	customerTransactionRepository := repository.NewCustomerTransactionRepository()
	customerTransactionService := service.NewCustomerTransactionService(
		customerTransactionRepository, customerTenorRepository, DB, validator)
	customerTransactionController := controller.NewCustomerTransactionController(customerTransactionService)

	// init router
	httpRouter := route.NewHttpRouter(customerController, customerTenorController, customerTransactionController)

	// run server
	serverAddress := fmt.Sprintf("localhost:%s", utils.ViperEnv("PORT"))
	server := http.Server{
		Handler: httpRouter,
		Addr:    serverAddress,
	}
	err := server.ListenAndServe()
	exception.PanicIfError(err)
}
