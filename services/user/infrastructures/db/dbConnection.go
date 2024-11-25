package db

import (
	"architecture_template/constants/notis"
	"architecture_template/services/user/entities"
	"database/sql"
	"errors"
	"log"
	"os"
)

func ConnectDB() (*sql.DB, error) {
	cnn, err := sql.Open(entities.GetDatabaseServer(), os.Getenv("CNN_STR"))

	if err != nil {
		log.Print("Fail to get access to database")
		return nil, errors.New(notis.InternalErr)
	}

	return cnn, nil
}
