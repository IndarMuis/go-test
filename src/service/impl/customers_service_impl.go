package impl

import (
	"database/sql"
	"github.com/IndarMuis/pt-xyz.git/src/models/dto"
	"github.com/IndarMuis/pt-xyz.git/src/repository"
	"github.com/IndarMuis/pt-xyz.git/src/service"
	"github.com/go-playground/validator"
)

type CustomerServiceImpl struct {
	CustomerRepository repository.CustomerRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func (customer *CustomerServiceImpl) CreateCustomer(customers dto.Customers) dto.Customers {
	//TODO implement me
	panic("implement me")
}

func NewCustomerService(CustomerRepository repository.CustomerRepository,
	DB *sql.DB,
	Validate *validator.Validate) service.CustomerService {
	return &CustomerServiceImpl{
		CustomerRepository: CustomerRepository,
		DB:                 DB,
		Validate:           Validate,
	}
}
