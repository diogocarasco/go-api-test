package main

import (
	configs "github.com/diogocarasco/go-api-test/config"
	"github.com/diogocarasco/go-api-test/database"
	"github.com/diogocarasco/go-api-test/utils"
)

func main() {

	utils.StartupMessage()

	database.InitializeDB()

	server := configs.GetServer()

	server.Run(":8080")

}
