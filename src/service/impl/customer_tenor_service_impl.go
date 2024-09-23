package impl

import (
	"context"
	"database/sql"
	"github.com/IndarMuis/pt-xyz.git/src/common/exception"
	"github.com/IndarMuis/pt-xyz.git/src/common/utils"
	"github.com/IndarMuis/pt-xyz.git/src/models/dto"
	"github.com/IndarMuis/pt-xyz.git/src/models/entity"
	"github.com/IndarMuis/pt-xyz.git/src/repository"
	"github.com/IndarMuis/pt-xyz.git/src/service"
	"github.com/go-playground/validator"
)

type CustomerTenorServiceImpl struct {
	CustomerTenorRepository repository.CustomerTenorRepository
	DB                      *sql.DB
	Validate                *validator.Validate
}

func (service *CustomerTenorServiceImpl) UpdateAmountUsed(ctx context.Context, request dto.UpdateCustomerTenorDto) *dto.UpdateCustomerTenorDto {
	err := service.Validate.Struct(&request)
	exception.PanicIfError(err)

	tx, err := service.DB.Begin()
	exception.PanicIfError(err)
	defer utils.CommitOrRollback(tx)

	customerTenor := entity.CustomersTenor{
		CustomerId:       request.CustomerId,
		TenorAmountLimit: request.TenorAmountLimit,
	}
	service.CustomerTenorRepository.UpdateAmountUsed(ctx, tx, customerTenor)

	return &request
}

func (service *CustomerTenorServiceImpl) Create(ctx context.Context, request dto.CreateCustomerTenorDto) *dto.CreateCustomerTenorDto {
	err := service.Validate.Struct(&request)
	exception.PanicIfError(err)

	tx, err := service.DB.Begin()
	exception.PanicIfError(err)
	defer utils.CommitOrRollback(tx)

	customerTenor := entity.CustomersTenor{
		CustomerId:       request.CustomerId,
		TotalMonth:       request.TotalMonth,
		TenorAmountLimit: request.TenorAmountLimit,
		StartDate:        request.StartDate,
		EndDate:          request.EndDate,
	}
	service.CustomerTenorRepository.Create(ctx, tx, customerTenor)

	return &request
}

func (service *CustomerTenorServiceImpl) SetCustomerTenor(ctx context.Context, request dto.ActivateTenorDto) *dto.ActivateTenorDto {
	err := service.Validate.Struct(&request)
	exception.PanicIfError(err)

	tx, err := service.DB.Begin()
	exception.PanicIfError(err)
	defer utils.CommitOrRollback(tx)

	customerTenor := entity.CustomersTenor{
		CustomerId: request.CustomerId,
		TotalMonth: request.TotalMonth,
		Status:     request.Status,
	}
	service.CustomerTenorRepository.SetCustomerTenor(ctx, tx, customerTenor)

	return &request

}

func (service *CustomerTenorServiceImpl) FindTenorByCustomerId(ctx context.Context, request dto.FindTenorByCustomerIdDto) *[]dto.CustomerTenorDto {
	tx, err := service.DB.Begin()
	exception.PanicIfError(err)
	defer utils.CommitOrRollback(tx)

	results := service.CustomerTenorRepository.FindTenorByCustomerId(ctx, tx, request.CustomerId)
	if len(results) < 1 {
		panic(exception.NewNotFoundError("not found"))
	}

	var customersTenor []dto.CustomerTenorDto
	for _, result := range results {
		customersTenor = append(customersTenor, dto.CustomerTenorDto{
			CustomerId:       result.CustomerId,
			TotalMonth:       result.TotalMonth,
			TenorAmountLimit: result.TenorAmountLimit,
			AmountUsed:       result.AmountUsed,
			StartDate:        result.StartDate,
			EndDate:          result.EndDate,
			Status:           result.Status,
		})
	}

	return &customersTenor
}

func NewCustomerTenorService(customerTenorRepository repository.CustomerTenorRepository,
	db *sql.DB, validate *validator.Validate) service.CustomerTenorService {
	return &CustomerTenorServiceImpl{
		CustomerTenorRepository: customerTenorRepository,
		DB:                      db,
		Validate:                validate,
	}
}
