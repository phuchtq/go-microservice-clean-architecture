package cmd

import (
	"architecture_template/constants/notis"
	"architecture_template/services/role/infrastructures/routes"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

func setUp() {
	// Start gin engine - api
	routes.InitializeAPIRoutes()

	// Start gRPC service
	routes.InitializeGRPCRoute()
}

// Config command
var rootCmd = &cobra.Command{
	Use:     "role-service",                                  // Command to start service
	Short:   "Service supports request about roles.",         // Short description about command
	Aliases: []string{"command 1", "command 2", "command 3"}, // alternative commands
	Run: func(cmd *cobra.Command, args []string) { // Start command
		log.Println("Run role service.")

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
		log.Fatalln(fmt.Sprintf(notis.CmdExecuteErrMsg, "Role") + err.Error())
	}
}
