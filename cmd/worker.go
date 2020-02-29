package cmd

import (
	"github.com/spf13/cobra"
	"kobe/pkg/worker"
)

var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		worker.Run()
	},
}

func init() {
	rootCmd.AddCommand(workerCmd)

}
