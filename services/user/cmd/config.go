package cmd

import (
	"architecture_template/constants/notis"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func config() {
	// Load env file
	if err := godotenv.Load(); err != nil {
		log.Fatal(fmt.Sprintf(notis.EnvLoadErr, "User") + err.Error())
	}

	// More configuration
}
