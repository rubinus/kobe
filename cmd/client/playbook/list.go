package playbook

import "github.com/spf13/cobra"

var playbookListCmd = &cobra.Command{
	Use: "list",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	Cmd.AddCommand(playbookListCmd)
}
