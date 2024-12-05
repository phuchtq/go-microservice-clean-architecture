package cmd

import (
	"architecture_template/constants/migrations"
	"architecture_template/helper/db"
	envvar "architecture_template/services/role/constants/envVar"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func migrate() {
	if err := db.Migrate(
		os.Getenv(envvar.DbCnnStr), // connection string
		"User",
		migrations.MigrateRequest, // Command
	); err != nil {
		fmt.Println(err.Error())
	}
}

var migrateCmd = &cobra.Command{
	Use:     "user-service migrate",                                   // Command to start service
	Short:   "Command to execute database migration in user service.", // Short description about command
	Aliases: []string{"command 1", "command 2", "command 3"},          // alternative commands
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Execute databse migration in user service.")
		migrate()
	},
}
