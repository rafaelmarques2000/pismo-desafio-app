package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"pismo-desafio-app/pkg/model"
)

var dbConnection *gorm.DB

func OpenConnection() (*gorm.DB, error) {
	if dbConnection == nil {
		connection, err := gorm.Open(sqlite.Open("./config/database/pismo_app.db"), &gorm.Config{
			SkipDefaultTransaction: true,
			PrepareStmt:            true,
		})
		dbConnection = connection
		return dbConnection, err
	}
	return dbConnection, nil
}

func AutoMigratesDatabase() {
	migrateError := dbConnection.AutoMigrate(
		&model.Account{},
		&model.OperationType{},
		&model.Transaction{},
	)
	if migrateError != nil {
		panic("Falha ao efetuar migração de banco de dados")
	}
}
