package exception

import (
	"github.com/IndarMuis/pt-xyz.git/src/common/utils"
	"github.com/IndarMuis/pt-xyz.git/src/models/dto"
	"github.com/go-playground/validator"
	"net/http"
)

func ExceptionHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {

	if notFoundError(writer, request, err) {
		return
	}

	if validationErrors(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)
}

func validationErrors(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		webResponse := &dto.ResponseTemplate{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   exception.Error(),
		}

		utils.WriteToResponseBody(writer, webResponse)

		return true
	} else {
		return false
	}
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		webResponse := dto.ResponseTemplate{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   exception.Error,
		}

		utils.WriteToResponseBody(writer, &webResponse)
		return true
	} else {
		return false
	}
}

//func generalError(writer http.ResponseWriter, request *http.Request, err interface{}) {
//	writer.Header().Add("Content-Type", "application/json")
//}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := dto.ResponseTemplate{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   err,
	}

	utils.WriteToResponseBody(writer, &webResponse)

}
