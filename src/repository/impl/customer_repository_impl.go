package impl

import (
	"database/sql"
	"github.com/IndarMuis/pt-xyz.git/src/models/dto"
	"github.com/IndarMuis/pt-xyz.git/src/repository"
)

func NewCustomerRepository(DB *sql.DB) repository.CustomerRepository {
	return &CustomerRepositoryImpl{DB: DB}
}

type CustomerRepositoryImpl struct {
	*sql.DB
}

func (customer *CustomerRepositoryImpl) CreateCustomer(customers dto.Customers) {
	//TODO implement me
	panic("implement me")
}

func (customer *CustomerRepositoryImpl) UpdateCustomer(customers dto.Customers) {
	//TODO implement me
	panic("implement me")
}
