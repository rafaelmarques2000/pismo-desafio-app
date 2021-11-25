package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"pismo-desafio-app/pkg/model"
)

var dbConnection *gorm.DB

func OpenConnection() (*gorm.DB, error) {
	connection, err := gorm.Open(sqlite.Open("./config/database/pismo_app.db"), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	dbConnection = connection
	return dbConnection, err
}

func AutoMigratesDatabase() {
	migrateAccountError := dbConnection.AutoMigrate(&model.Account{})
	if migrateAccountError != nil {
		return
	}
	migrateOperationTypeError := dbConnection.AutoMigrate(&model.OperationType{})
	if migrateOperationTypeError != nil {
		return
	}
	migrateTransactionError := dbConnection.AutoMigrate(&model.Transaction{})
	if migrateTransactionError != nil {
		return
	}
}
