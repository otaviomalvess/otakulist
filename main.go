package main

import (
	"otavio-alves/OtakuList/database"
	"otavio-alves/OtakuList/service"
)

// Starts application
func main() {

	// Call DB client constructor
	database.CreateClient()

	// Call start server function
	service.StartServer()
}
