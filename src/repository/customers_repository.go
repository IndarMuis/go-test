package repository

import (
	"github.com/IndarMuis/pt-xyz.git/src/models/dto"
)

type CustomerRepository interface {
	CreateCustomer(customers dto.Customers)
	UpdateCustomer(customers dto.Customers)
}
