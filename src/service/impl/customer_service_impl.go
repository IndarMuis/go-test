package impl

import (
	"context"
	"database/sql"
	"errors"
	"github.com/IndarMuis/pt-xyz.git/src/common/exception"
	"github.com/IndarMuis/pt-xyz.git/src/common/utils"
	"github.com/IndarMuis/pt-xyz.git/src/models/dto"
	"github.com/IndarMuis/pt-xyz.git/src/models/entity"
	"github.com/IndarMuis/pt-xyz.git/src/repository"
	"github.com/IndarMuis/pt-xyz.git/src/service"
	"github.com/go-playground/validator"
)

type CustomerServiceImpl struct {
	CustomerRepository repository.CustomerRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func (service *CustomerServiceImpl) Create(ctx context.Context, request dto.CreateCustomerDto) (*dto.CreateCustomerDto, error) {
	err := service.Validate.Struct(request)
	exception.PanicIfError(err)

	tx, err := service.DB.Begin()
	exception.PanicIfError(err)
	defer utils.CommitOrRollback(tx)

	customer := entity.Customers{
		Nik:         request.Nik,
		FullName:    request.FullName,
		LegalName:   request.LegalName,
		BirthPlace:  request.BirthPlace,
		BirthDate:   request.BirthDate,
		Salary:      request.Salary,
		KTPPhoto:    request.KTPPhoto,
		SelfiePhoto: request.SelfiePhoto,
	}

	createCustomer := service.CustomerRepository.Create(ctx, tx, customer)
	if createCustomer.Id == 0 {
		return nil, errors.New("failed to create customer")
	} else {
		return &request, nil
	}

}

func (service *CustomerServiceImpl) FindAll(ctx context.Context) []dto.CustomerDto {
	tx, err := service.DB.Begin()
	exception.PanicIfError(err)
	defer utils.CommitOrRollback(tx)

	results := service.CustomerRepository.FindAll(ctx, tx)

	var customers []dto.CustomerDto

	for _, customer := range results {
		customers = append(customers, dto.CustomerDto{
			Nik:         customer.Nik,
			FullName:    customer.FullName,
			LegalName:   customer.LegalName,
			BirthPlace:  customer.BirthPlace,
			BirthDate:   customer.BirthDate,
			Salary:      customer.Salary,
			KTPPhoto:    customer.KTPPhoto,
			SelfiePhoto: customer.SelfiePhoto,
		})
	}

	return customers
}

func NewCustomerService(CustomerRepository repository.CustomerRepository,
	DB *sql.DB,
	Validate *validator.Validate) service.CustomerService {
	return &CustomerServiceImpl{
		CustomerRepository: CustomerRepository,
		Validate:           Validate,
		DB:                 DB,
	}
}
