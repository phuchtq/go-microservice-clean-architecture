package cmd

import (
	"architecture_template/constants/notis"
	"architecture_template/services/user/infrastructures/routes"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

func setUp() {
	// Start gin engine - api
	routes.InitializeAPIRoutes()
}

// Config command
var rootCmd = &cobra.Command{
	Use:     "user-service",                                  // Command to start service
	Short:   "Service supports request about accounts.",      // Short description about command
	Aliases: []string{"command 1", "command 2", "command 3"}, // alternative commands
	Run: func(cmd *cobra.Command, args []string) { // Start command
		log.Println("Run user service.")

		// Load configuration
		config()

		// Set up service
		setUp()
	},
}

func Execute() {
	// Add migrate/rollback database command (optional)
	rootCmd.AddCommand(migrateCmd, rollbackCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(fmt.Sprintf(notis.CmdExecuteErrMsg, "User") + err.Error())
	}
}
