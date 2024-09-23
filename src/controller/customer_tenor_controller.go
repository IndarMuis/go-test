package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type CustomerTenorController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	SetCustomerTenor(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdateAmountUsed(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindTenorByCustomerId(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
