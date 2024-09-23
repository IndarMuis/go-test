package repository

import (
	"context"
	"database/sql"
	"github.com/IndarMuis/pt-xyz.git/src/models/entity"
)

type CustomerTransactionRepository interface {
	Create(ctx context.Context, tx *sql.Tx, transaction entity.CustomersTransaction) entity.CustomersTransaction
	FindCustomerTransaction(ctx context.Context, tx *sql.Tx, customerId int) []entity.CustomersTransaction
}
