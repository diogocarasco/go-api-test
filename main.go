package main

import (
	configs "github.com/diogocarasco/go-api-test/config"
	"github.com/diogocarasco/go-api-test/database"
)

func main() {

	database.InitializeDB()

	server := configs.GetServer()

	server.Run(":8080")

}
