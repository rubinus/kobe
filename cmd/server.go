package cmd

import (
	"github.com/spf13/cobra"
	"kobe/pkg/server"
	"log"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if err := server.Run(); err != nil {
			log.Panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
