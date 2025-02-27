package db

import (
	"architecture_template/constants/notis"
	envvar "architecture_template/services/role/constants/envVar"
	"architecture_template/services/role/entities"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

const (
	backUpDbServer string = "Your back up database server"

	backUpDbCnnStr string = "Your back up database connection string"
)

func ConnectDB() (*sql.DB, error) {
	var logger = &log.Logger{}
	var service string = "Role"

	var dbServer string = entities.GetDatabaseServer()
	if dbServer == "" {
		logger.Println(fmt.Sprintf(notis.DbServerNotSetMsg, service))
		dbServer = backUpDbServer
	}

	var cnnStr string = os.Getenv(envvar.DbCnnStr)
	if cnnStr == "" {
		logger.Println(fmt.Sprintf(notis.DbServerNotSetMsg, service))

		if err := os.Setenv(envvar.DbCnnStr, backUpDbCnnStr); err != nil {
			logger.Println(fmt.Sprintf(notis.DbSetConnectionStrErrMsg, service) + err.Error())
		}

		cnnStr = backUpDbCnnStr
	}

	cnn, err := sql.Open(dbServer, cnnStr)

	if err != nil {
		logger.Println(fmt.Sprintf(notis.DbConnectionMsg, service) + err.Error())
		return nil, errors.New(notis.InternalErr)
	}

	return cnn, nil
}
