package playbook

import "github.com/spf13/cobra"

var playbookRunCmd = &cobra.Command{
	Use: "run",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	Cmd.AddCommand(playbookRunCmd)
}