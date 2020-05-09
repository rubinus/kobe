package playbook

import (
	"github.com/spf13/cobra"
	"kobe/cmd/client/root"
)

var Cmd = &cobra.Command{
	Use: "playbook",
}

func init() {
	root.Cmd.AddCommand(Cmd)
}

