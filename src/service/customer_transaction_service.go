package service

import (
	"context"
	"github.com/IndarMuis/pt-xyz.git/src/models/dto"
)

type CustomerTransactionService interface {
	Create(ctx context.Context, request dto.CreateCustomerTransactionDto) *dto.CreateCustomerTransactionDto
	FindCustomerTransaction(ctx context.Context, customerId int) *[]dto.CustomerTransactionDto
}
