package project

import "github.com/spf13/cobra"

var projectCreateCmd = &cobra.Command{
	Use: "create",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	Cmd.AddCommand(projectCreateCmd)
}
