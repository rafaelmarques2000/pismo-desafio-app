package main

import (
	"pismo-desafio-app/config/database"
	"pismo-desafio-app/config/database/seed/operation_type_seeder"
	"pismo-desafio-app/server"
)

func main()  {
	setupDatabase()
	server.Start()
}

func setupDatabase() {
	_, err := database.OpenConnection()
	if err != nil {
		panic("Falha ao iniciar app não foi possivel conectar com o banco de dados")
	}
	database.AutoMigratesDatabase()
	operation_type_seeder.RunSeeder()
}
