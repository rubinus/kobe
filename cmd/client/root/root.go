package root

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "kobe",
	Short: "A kobe client cli",
}

func init() {
	Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
