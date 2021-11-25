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
	dbConnection.AutoMigrate(&model.Account{})
	dbConnection.AutoMigrate(&model.OperationType{})
	dbConnection.AutoMigrate(&model.Transaction{})
}
