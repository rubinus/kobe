package root

import (
	"github.com/spf13/cobra"
	"github.com/KubeOperator/kobe/cmd/client/playbook"
	"github.com/KubeOperator/kobe/cmd/client/project"
	"github.com/KubeOperator/kobe/cmd/client/task"
)

var Cmd = &cobra.Command{
	Use:   "kobe",
	Short: "A kobe client cli",
}

func init() {
	Cmd.AddCommand(project.Cmd)
	Cmd.AddCommand(playbook.Cmd)
	Cmd.AddCommand(task.Cmd)
}
