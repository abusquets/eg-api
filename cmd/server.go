package cmd

import (
	"eventsguard/internal/app"
	"eventsguard/internal/di"
	"eventsguard/internal/infrastructure/server"

	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Server application",
	Long:  "HTTP Server application",
	Run: func(cmd *cobra.Command, args []string) {
		app := fx.New(
			di.BaseModule,
			app.Module,
			server.Module,
		)

		app.Run()
	},
}
