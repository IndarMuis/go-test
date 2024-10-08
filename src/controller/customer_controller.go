package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type CustomerController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
