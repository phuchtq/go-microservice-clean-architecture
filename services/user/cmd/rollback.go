package cmd

import (
	"architecture_template/constants/migrations"
	"architecture_template/helper/db"
	envvar "architecture_template/services/user/constants/envVar"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func rollback() {
	if err := db.Migrate(
		os.Getenv(envvar.DbCnnStr), // Connection string
		"User",                     // Service
		migrations.RollbackRequest, // Request command
	); err != nil {
		fmt.Println(err.Error())
	}
}

var rollbackCmd = &cobra.Command{
	Use:     "user-service rollback",                                 // Command to start service
	Short:   "Command to execute database rollback in user service.", // Short description about command
	Aliases: []string{"command 1", "command 2", "command 3"},         // alternative commands
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Execute database rollback in user service.")
		rollback()
	},
}
