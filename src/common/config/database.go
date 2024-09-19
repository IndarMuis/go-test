package config

import (
	"database/sql"
	"fmt"
	"github.com/IndarMuis/pt-xyz.git/src/common/exception"
	"github.com/IndarMuis/pt-xyz.git/src/common/utils"
	"time"
)

func NewDatabase() *sql.DB {
	databaseHost := utils.ViperEnv("DATABASE_HOST")
	databasePort := utils.ViperEnv("DATABASE_PORT")
	databaseUser := utils.ViperEnv("DATABASE_USER")
	databasePassword := utils.ViperEnv("DATABASE_PASSWORD")
	databaseName := utils.ViperEnv("DATABASE_NAME")

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		databaseUser, databasePassword, databaseHost,
		databasePort, databaseName,
	)

	db, err := sql.Open("mysql", dataSourceName)
	exception.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
