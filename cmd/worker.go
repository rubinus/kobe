package cmd

import (
	"github.com/spf13/cobra"
	container "kobe/pkg/worker"
)

var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		container.Run()
	},
}

func init() {
	rootCmd.AddCommand(workerCmd)

}
