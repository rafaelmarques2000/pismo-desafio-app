package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"pismo-desafio-app/app/model/account"
	"pismo-desafio-app/app/model/operationtype"
	"pismo-desafio-app/app/model/transaction"
)

var dbConnection *gorm.DB

func OpenConnection() (*gorm.DB,error) {
	connection, err := gorm.Open(sqlite.Open("pismo_app.db"), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt: true,
	})
	dbConnection = connection
	return dbConnection,err
}

func AutoMigratesDatabase()  {
	dbConnection.AutoMigrate(&account.Account{})
	dbConnection.AutoMigrate(&operationtype.OperationType{})
	dbConnection.AutoMigrate(&transaction.Transaction{})
}

