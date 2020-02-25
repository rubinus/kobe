package cmd

import (
	"github.com/spf13/cobra"
	"kobe/pkg/server"
	"log"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run a http server",
	Run: func(cmd *cobra.Command, args []string) {
		if err := server.RunServer(); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
