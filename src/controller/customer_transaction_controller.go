package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type CustomerTransactionController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindCustomerTransaction(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
