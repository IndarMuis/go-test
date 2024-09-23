package impl

import (
	"github.com/IndarMuis/pt-xyz.git/src/common/exception"
	"github.com/IndarMuis/pt-xyz.git/src/common/utils"
	"github.com/IndarMuis/pt-xyz.git/src/controller"
	"github.com/IndarMuis/pt-xyz.git/src/models/dto"
	"github.com/IndarMuis/pt-xyz.git/src/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type CustomerControllerImpl struct {
	customerService service.CustomerService
}

func (controller *CustomerControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	createCustomerDto := dto.CreateCustomerDto{}
	utils.ReadFromRequestBody(request, &createCustomerDto)

	response, err := controller.customerService.Create(request.Context(), createCustomerDto)
	exception.PanicIfError(err)

	webResponse := dto.ResponseTemplate{
		Code:   201,
		Status: "Created",
		Data:   response,
	}
	utils.WriteToResponseBody(writer, webResponse)
}

func (controller *CustomerControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	response := controller.customerService.FindAll(request.Context())

	webResponse := dto.ResponseTemplate{
		Code:   200,
		Status: "OK",
		Data:   response,
	}
	utils.WriteToResponseBody(writer, webResponse)
}

func NewCustomerController(customerService service.CustomerService) controller.CustomerController {
	return &CustomerControllerImpl{customerService: customerService}
}
