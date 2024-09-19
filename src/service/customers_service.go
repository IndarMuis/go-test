package service

import "github.com/IndarMuis/pt-xyz.git/src/models/dto"

type CustomerService interface {
	CreateCustomer(customers dto.Customers) dto.Customers
}
