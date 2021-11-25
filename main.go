package main

import (
	"os"
	"pismo-desafio-app/config/database"
	"pismo-desafio-app/config/database/seed"
	"pismo-desafio-app/server"
)

func main() {
	setupDatabase()
	registerEnvVariables()
	server.NewServer().Start()
}

func registerEnvVariables()  {
	err := os.Setenv("SERVER_PORT", ":8080")
	if err != nil {
		return
	}
}

func setupDatabase() {
	_, err := database.OpenConnection()
	if err != nil {
		panic("Falha ao iniciar app não foi possivel conectar com o banco de dados")
	}
	database.AutoMigratesDatabase()
	seed.NewOperationTypeSeeder().RunSeeder()
}
