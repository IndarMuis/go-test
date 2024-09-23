package repository

import (
	"context"
	"database/sql"
	"github.com/IndarMuis/pt-xyz.git/src/models/entity"
)

type CustomerTenorRepository interface {
	Create(ctx context.Context, tx *sql.Tx, tenor entity.CustomersTenor) entity.CustomersTenor
	SetCustomerTenor(ctx context.Context, tx *sql.Tx, tenor entity.CustomersTenor) entity.CustomersTenor
	UpdateAmountUsed(ctx context.Context, tx *sql.Tx, tenor entity.CustomersTenor) entity.CustomersTenor
	FindTenorByCustomerId(ctx context.Context, tx *sql.Tx, customerId int) []entity.CustomersTenor
}
