package impl

import (
	"context"
	"database/sql"
	"github.com/IndarMuis/pt-xyz.git/src/common/exception"
	"github.com/IndarMuis/pt-xyz.git/src/models/entity"
	"github.com/IndarMuis/pt-xyz.git/src/repository"
)

type CustomerTransactionRepositoryImpl struct {
}

func (repository *CustomerTransactionRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, transaction entity.CustomersTransaction) entity.CustomersTransaction {
	query := "INSERT INTO customers_transaction(contract_number, customer_id, otr, admin_fee, installment_amount, interest_amount, asset_name)" +
		" VALUES (?, ?, ?, ?, ?, ?, ?)"
	_, err := tx.ExecContext(ctx, query, transaction.ContractNumber, transaction.CustomerId, transaction.OTR, transaction.AdminFee,
		transaction.InstallmentAmount, transaction.InterestAmount, transaction.AssetName)
	exception.PanicIfError(err)

	return transaction

}

func (repository *CustomerTransactionRepositoryImpl) FindCustomerTransaction(ctx context.Context, tx *sql.Tx, customerId int) []entity.CustomersTransaction {
	query := "SELECT contract_number, customer_id, otr, admin_fee, " +
		"installment_amount, interest_amount, asset_name, transaction_date " +
		"FROM customers_transaction WHERE customer_id = ?"
	rows, err := tx.QueryContext(ctx, query, customerId)
	exception.PanicIfError(err)

	var customerTransactions []entity.CustomersTransaction
	for rows.Next() {
		transaction := entity.CustomersTransaction{}
		err := rows.Scan(&transaction.ContractNumber, &transaction.CustomerId, &transaction.OTR, &transaction.AdminFee,
			&transaction.InstallmentAmount, &transaction.InterestAmount, &transaction.AssetName, &transaction.TransactionDate)
		exception.PanicIfError(err)
		customerTransactions = append(customerTransactions, transaction)
	}
	return customerTransactions
}

func NewCustomerTransactionRepository() repository.CustomerTransactionRepository {
	return &CustomerTransactionRepositoryImpl{}
}
