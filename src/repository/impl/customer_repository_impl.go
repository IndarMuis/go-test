package impl

import (
	"context"
	"database/sql"
	"github.com/IndarMuis/pt-xyz.git/src/common/exception"
	"github.com/IndarMuis/pt-xyz.git/src/models/entity"
	"github.com/IndarMuis/pt-xyz.git/src/repository"
)

func NewCustomerRepository(DB *sql.DB) repository.CustomerRepository {
	return &CustomerRepositoryImpl{DB: DB}
}

type CustomerRepositoryImpl struct {
	*sql.DB
}

func (repository *CustomerRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, customer entity.Customers) entity.Customers {
	query := "INSERT INTO customers(nik, full_name, legal_name, birth_place, birth_date, salary, ktp_photo, selfie_photo)" +
		" VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, query, customer.Nik, customer.FullName, customer.LegalName,
		customer.BirthPlace, customer.BirthDate, customer.Salary, customer.KTPPhoto, customer.SelfiePhoto)
	exception.PanicIfError(err)

	id, err := result.LastInsertId()
	exception.PanicIfError(err)

	customer.Id = int(id)

	return customer
}

func (repository *CustomerRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, customer entity.Customers) entity.Customers {
	//TODO implement me
	panic("implement me")
}

func (repository *CustomerRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []entity.Customers {
	//TODO implement me
	panic("implement me")
}
