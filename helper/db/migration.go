package db

import (
	"architecture_template/constants/migrations"
	"architecture_template/constants/notis"
	"errors"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
)

const (
	migrationDir string = "file://infrastructures/db/migrations" // Shared structure through services
)

func Migrate(cnn, service, action string) error {
	var logger = &log.Logger{}
	var errPrefixMsg string = fmt.Sprintf(notis.DbMigrationErrMsg, service)

	migration, err := migrate.New(
		migrationDir,
		cnn,
	)

	if err != nil {
		logger.Println(errPrefixMsg + err.Error())
		return errors.New(notis.DbMigrationInformMsg)
	}

	var res error

	switch action {
	case migrations.MigrateRequest:
		res = migration.Up()
	case migrations.RollbackRequest:
		res = migration.Down()
	default: // No method found
		return errors.New("Invalid migration command.")
	}

	if res != nil && res != migrate.ErrNoChange {
		logger.Println(errPrefixMsg + err.Error())
		return errors.New(notis.DbMigrationInformMsg)
	}

	logger.Println(action + "successful.")
	return nil
}
