package main

import (
	configs "github.com/diogocarasco/go-api-test/config"
	"github.com/diogocarasco/go-api-test/database"
	"github.com/diogocarasco/go-api-test/utils"
	// gin-swagger middleware
	// swagger embed files
)

func main() {

	utils.StartupMessage()

	database.InitializeDB("live")

	server := configs.GetServer()

	server.Run(":8080")

}
