package controller

import (
	"github.com/IndarMuis/pt-xyz.git/src/common/exception"
	"github.com/julienschmidt/httprouter"
)

func NewHttpRouter(customerController CustomerController,
	customerTenorController CustomerTenorController,
	customerTransactionController CustomerTransactionController) *httprouter.Router {
	router := httprouter.New()

	// customer router
	router.GET("/api/v1/customers", customerController.FindAll)
	router.POST("/api/v1/customers", customerController.Create)

	// customer tenor router
	router.POST("/api/v1/customers_tenor", customerTenorController.Create)
	router.PUT("/api/v1/customers_tenor/set", customerTenorController.SetCustomerTenor)
	router.PUT("/api/v1/customers_tenor/amount_used", customerTenorController.UpdateAmountUsed)
	router.GET("/api/v1/customers_tenor/:customerId", customerTenorController.FindTenorByCustomerId)

	// customer transaction router
	router.POST("/api/v1/customers_transaction", customerTransactionController.Create)
	router.GET("/api/v1/customers_transaction/:customerId", customerTransactionController.FindCustomerTransaction)
	router.PanicHandler = exception.ExceptionHandler

	return router
}
