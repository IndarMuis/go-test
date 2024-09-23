package impl

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/IndarMuis/pt-xyz.git/src/common/exception"
	"github.com/IndarMuis/pt-xyz.git/src/common/utils"
	"github.com/IndarMuis/pt-xyz.git/src/models/dto"
	"github.com/IndarMuis/pt-xyz.git/src/models/entity"
	"github.com/IndarMuis/pt-xyz.git/src/repository"
	"github.com/IndarMuis/pt-xyz.git/src/service"
	"github.com/go-playground/validator"
	"math/rand"
	"time"
)

type CustomerTransactionServiceImpl struct {
	CustomerTransactionRepository repository.CustomerTransactionRepository
	CustomerTenorRepository       repository.CustomerTenorRepository
	DB                            *sql.DB
	Validate                      *validator.Validate
}

func (service *CustomerTransactionServiceImpl) Create(ctx context.Context, request dto.CreateCustomerTransactionDto) *dto.CreateCustomerTransactionDto {
	err := service.Validate.Struct(&request)
	exception.PanicIfError(err)

	tx, err := service.DB.Begin()
	exception.PanicIfError(err)
	defer utils.CommitOrRollback(tx)

	rand.Seed(time.Now().UnixNano())
	generateContractNumber := rand.Intn(90) + 50
	request.ContractNumber = fmt.Sprintf("XPRODUCT-%d", generateContractNumber)

	customerTransaction := entity.CustomersTransaction{
		ContractNumber:    request.ContractNumber,
		CustomerId:        request.CustomerId,
		OTR:               request.Otr,
		AdminFee:          request.AdminFee,
		InstallmentAmount: request.InstallmentAmount,
		InterestAmount:    request.InterestAmount,
		AssetName:         request.AssetName,
	}
	customerTenor := entity.CustomersTenor{
		CustomerId: request.CustomerId,
		AmountUsed: request.Otr,
	}

	// concurrent process
	go service.CustomerTransactionRepository.Create(ctx, tx, customerTransaction)
	go service.CustomerTenorRepository.UpdateAmountUsed(ctx, tx, customerTenor)

	return &request
}

func (service *CustomerTransactionServiceImpl) FindCustomerTransaction(ctx context.Context, customerId int) *[]dto.CustomerTransactionDto {
	tx, err := service.DB.Begin()
	exception.PanicIfError(err)
	defer utils.CommitOrRollback(tx)

	var customerTransaction []dto.CustomerTransactionDto
	results := service.CustomerTransactionRepository.FindCustomerTransaction(ctx, tx, customerId)
	for _, result := range results {
		customerTransaction = append(customerTransaction, dto.CustomerTransactionDto{
			ContractNumber:    result.ContractNumber,
			CustomerId:        result.CustomerId,
			Otr:               result.OTR,
			AdminFee:          result.AdminFee,
			InstallmentAmount: result.InstallmentAmount,
			InterestAmount:    result.InterestAmount,
			AssetName:         result.AssetName,
		})
	}

	return &customerTransaction
}

func NewCustomerTransactionService(
	customerTransactionRepository repository.CustomerTransactionRepository, customerTenorRepository repository.CustomerTenorRepository,
	db *sql.DB, validate *validator.Validate) service.CustomerTransactionService {
	return &CustomerTransactionServiceImpl{
		CustomerTransactionRepository: customerTransactionRepository,
		CustomerTenorRepository:       customerTenorRepository,
		DB:                            db,
		Validate:                      validate,
	}
}
