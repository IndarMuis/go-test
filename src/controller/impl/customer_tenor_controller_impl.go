package impl

import (
	"github.com/IndarMuis/pt-xyz.git/src/common/exception"
	"github.com/IndarMuis/pt-xyz.git/src/common/utils"
	"github.com/IndarMuis/pt-xyz.git/src/controller"
	"github.com/IndarMuis/pt-xyz.git/src/models/dto"
	"github.com/IndarMuis/pt-xyz.git/src/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type CustomerTenorControllerImpl struct {
	CustomerTenorService service.CustomerTenorService
}

func (controller *CustomerTenorControllerImpl) UpdateAmountUsed(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	updateCustomerTenor := dto.UpdateCustomerTenorDto{}
	utils.ReadFromRequestBody(request, &updateCustomerTenor)

	response := controller.CustomerTenorService.UpdateAmountUsed(request.Context(), updateCustomerTenor)
	webResponse := dto.ResponseTemplate{
		Code:   200,
		Status: "OK",
		Data:   &response,
	}
	utils.WriteToResponseBody(writer, webResponse)
}

func (controller *CustomerTenorControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	createCustomerTenorDto := dto.CreateCustomerTenorDto{}
	utils.ReadFromRequestBody(request, &createCustomerTenorDto)

	response := controller.CustomerTenorService.Create(request.Context(), createCustomerTenorDto)
	webResponse := dto.ResponseTemplate{
		Code:   201,
		Status: "Created",
		Data:   &response,
	}
	utils.WriteToResponseBody(writer, webResponse)
}

func (controller *CustomerTenorControllerImpl) SetCustomerTenor(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	activateTenorDto := dto.ActivateTenorDto{}
	utils.ReadFromRequestBody(request, &activateTenorDto)

	response := controller.CustomerTenorService.SetCustomerTenor(request.Context(), activateTenorDto)
	webResponse := dto.ResponseTemplate{
		Code:   200,
		Status: "OK",
		Data:   response,
	}
	utils.WriteToResponseBody(writer, webResponse)
}

func (controller *CustomerTenorControllerImpl) FindTenorByCustomerId(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	customerId, err := strconv.Atoi(params.ByName("customerId"))
	exception.PanicIfError(err)

	requestDto := dto.FindTenorByCustomerIdDto{CustomerId: customerId}
	response := controller.CustomerTenorService.FindTenorByCustomerId(request.Context(), requestDto)
	webResponse := dto.ResponseTemplate{
		Code:   200,
		Status: "OK",
		Data:   response,
	}
	utils.WriteToResponseBody(writer, webResponse)
}

func NewCustomerTenorController(customerTenorService service.CustomerTenorService) controller.CustomerTenorController {
	return &CustomerTenorControllerImpl{CustomerTenorService: customerTenorService}
}
