package main

import (
	role_cmd "architecture_template/services/role/cmd"
	user_cmd "architecture_template/services/user/cmd"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// Set up
	setUp()

	// Execute role service
	role_cmd.Execute()

	// Execute user service
	user_cmd.Execute()
}

func setUp() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file in main - ", err.Error())
	}
}
