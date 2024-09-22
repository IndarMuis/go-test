package repository

import (
	"context"
	"database/sql"
	"github.com/IndarMuis/pt-xyz.git/src/models/entity"
)

type CustomerRepository interface {
	Create(ctx context.Context, tx *sql.Tx, customer entity.Customers) entity.Customers
	Update(ctx context.Context, tx *sql.Tx, customer entity.Customers) entity.Customers
	FindAll(ctx context.Context, tx *sql.Tx) []entity.Customers
}
