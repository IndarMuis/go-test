package service

import (
	"context"
	"github.com/IndarMuis/pt-xyz.git/src/models/dto"
)

type CustomerTenorService interface {
	Create(ctx context.Context, request dto.CreateCustomerTenorDto) *dto.CreateCustomerTenorDto
	SetCustomerTenor(ctx context.Context, request dto.ActivateTenorDto) *dto.ActivateTenorDto
	UpdateAmountUsed(ctx context.Context, request dto.UpdateCustomerTenorDto) *dto.UpdateCustomerTenorDto
	FindTenorByCustomerId(ctx context.Context, request dto.FindTenorByCustomerIdDto) *[]dto.CustomerTenorDto
}
