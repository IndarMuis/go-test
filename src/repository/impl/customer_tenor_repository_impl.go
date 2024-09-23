package impl

import (
	"context"
	"database/sql"
	"github.com/IndarMuis/pt-xyz.git/src/common/exception"
	"github.com/IndarMuis/pt-xyz.git/src/models/entity"
	"github.com/IndarMuis/pt-xyz.git/src/repository"
)

type CustomerTenorRepositoryImpl struct {
}

func (repository *CustomerTenorRepositoryImpl) UpdateAmountUsed(ctx context.Context, tx *sql.Tx, tenor entity.CustomersTenor) entity.CustomersTenor {
	query := "UPDATE customers_tenor SET amount_used = ? WHERE  customer_id = ? AND STATUS = 'ACTIVE'"
	_, err := tx.ExecContext(ctx, query, tenor.AmountUsed, tenor.CustomerId)
	exception.PanicIfError(err)
	return tenor
}

func (repository *CustomerTenorRepositoryImpl) FindTenorByCustomerId(ctx context.Context, tx *sql.Tx, customerId int) []entity.CustomersTenor {
	query := "SELECT customer_id, total_month, tenor_amount_limit, amount_used, start_date, end_date, status FROM customers_tenor WHERE customer_id = ?"
	rows, err := tx.QueryContext(ctx, query, customerId)
	defer rows.Close()
	exception.PanicIfError(err)

	var customersTenor []entity.CustomersTenor
	for rows.Next() {
		tenor := entity.CustomersTenor{}
		err := rows.Scan(&tenor.CustomerId, &tenor.TotalMonth, &tenor.TenorAmountLimit,
			&tenor.AmountUsed, &tenor.StartDate, &tenor.EndDate, &tenor.Status)
		exception.PanicIfError(err)
		customersTenor = append(customersTenor, tenor)
	}

	return customersTenor
}

func (repository *CustomerTenorRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, tenor entity.CustomersTenor) entity.CustomersTenor {
	query := "INSERT INTO customers_tenor(customer_id, total_month, tenor_amount_limit, start_date, end_date) " +
		"VALUES (?, ?, ?, ?, ?)"

	result, err := tx.ExecContext(ctx, query, tenor.CustomerId, tenor.TotalMonth, tenor.TenorAmountLimit,
		tenor.StartDate, tenor.EndDate)
	exception.PanicIfError(err)

	lastInsertId, err := result.LastInsertId()
	exception.PanicIfError(err)
	tenor.Id = int(lastInsertId)

	return tenor
}

func (repository *CustomerTenorRepositoryImpl) SetCustomerTenor(ctx context.Context, tx *sql.Tx, tenor entity.CustomersTenor) entity.CustomersTenor {

	query := "UPDATE customers_tenor SET status=? WHERE customer_id=? AND total_month=?"
	_, err := tx.ExecContext(ctx, query, tenor.Status, tenor.CustomerId, tenor.TotalMonth)
	exception.PanicIfError(err)

	return tenor
}

func NewCustomerTenorRepository() repository.CustomerTenorRepository {
	return &CustomerTenorRepositoryImpl{}
}
