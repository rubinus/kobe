package project

import (
	"github.com/spf13/cobra"
	"kobe/cmd/client/root"
)

var Cmd = &cobra.Command{
	Use: "project",
}

func init() {
	root.Cmd.AddCommand(Cmd)
}
