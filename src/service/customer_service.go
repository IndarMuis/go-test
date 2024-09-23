package service

import (
	"context"
	"github.com/IndarMuis/pt-xyz.git/src/models/dto"
)

type CustomerService interface {
	Create(ctx context.Context, request dto.CreateCustomerDto) (*dto.CreateCustomerDto, error)
	FindAll(ctx context.Context) []dto.CustomerDto
}
