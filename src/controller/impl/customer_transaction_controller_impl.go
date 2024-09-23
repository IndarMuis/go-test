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

type CustomerTransactionControllerImpl struct {
	CustomerTransactionService service.CustomerTransactionService
}

func (controller *CustomerTransactionControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	createCustomerTransaction := dto.CreateCustomerTransactionDto{}
	utils.ReadFromRequestBody(request, &createCustomerTransaction)

	response := controller.CustomerTransactionService.Create(request.Context(), createCustomerTransaction)
	webResponse := dto.ResponseTemplate{
		Code:   201,
		Status: "Created",
		Data:   &response,
	}
	utils.WriteToResponseBody(writer, webResponse)
}

func (controller *CustomerTransactionControllerImpl) FindCustomerTransaction(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	customerId, err := strconv.Atoi(params.ByName("customerId"))
	exception.PanicIfError(err)

	response := controller.CustomerTransactionService.FindCustomerTransaction(request.Context(), customerId)
	webResponse := dto.ResponseTemplate{
		Code:   200,
		Status: "OK",
		Data:   response,
	}
	utils.WriteToResponseBody(writer, webResponse)
}

func NewCustomerTransactionController(customerTransactionService service.CustomerTransactionService) controller.CustomerTransactionController {
	return &CustomerTransactionControllerImpl{
		CustomerTransactionService: customerTransactionService,
	}
}
