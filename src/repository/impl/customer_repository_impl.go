package impl

import (
	"context"
	"database/sql"
	"github.com/IndarMuis/pt-xyz.git/src/common/exception"
	"github.com/IndarMuis/pt-xyz.git/src/models/entity"
	"github.com/IndarMuis/pt-xyz.git/src/repository"
)

func NewCustomerRepository() repository.CustomerRepository {
	return &CustomerRepositoryImpl{}
}

type CustomerRepositoryImpl struct {
}

func (repository *CustomerRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, customer entity.Customers) entity.Customers {
	// create customer
	query := "INSERT INTO customers(nik, full_name, legal_name, birth_place, birth_date, salary, ktp_photo, selfie_photo)" +
		" VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, query, customer.Nik, customer.FullName, customer.LegalName,
		customer.BirthPlace, customer.BirthDate, customer.Salary, customer.KTPPhoto, customer.SelfiePhoto)
	exception.PanicIfError(err)

	id, err := result.LastInsertId()
	exception.PanicIfError(err)
	customer.Id = int(id)

	return customer
}

func (repository *CustomerRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []entity.Customers {
	query := "SELECT id, nik, full_name, legal_name, birth_place, birth_date, salary, ktp_photo, selfie_photo"
	rows, err := tx.QueryContext(ctx, query)
	defer rows.Close()
	exception.PanicIfError(err)

	var customers []entity.Customers
	for rows.Next() {
		customer := entity.Customers{}
		err := rows.Scan(&customer.Id, customer.Nik, &customer.FullName, &customer.LegalName,
			&customer.BirthPlace, &customer.BirthDate, &customer.Salary, &customer.KTPPhoto, &customer.SelfiePhoto)
		exception.PanicIfError(err)
		customers = append(customers, customer)
	}

	return customers
}
