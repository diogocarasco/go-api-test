package main

import (
	configs "go-api-test/config"
	"go-api-test/database"
)

func main() {

	database.InitializeDB()

	server := configs.GetServer()

	server.Run(":8080")

}
