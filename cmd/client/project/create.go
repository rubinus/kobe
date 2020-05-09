package project

import "github.com/spf13/cobra"

var projectListCmd = &cobra.Command{
	Use: "list",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	Cmd.AddCommand(projectListCmd)
}
