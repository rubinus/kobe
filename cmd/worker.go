package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"kobe/pkg/models"
)

var id string

var workerCmd = &cobra.Command{
	Use:   "inventory",
	Short: "dynamic inventory provider",
	Run: func(cmd *cobra.Command, args []string) {
		s, err := models.GetInventoryFromCache(id)
		if err != nil {
			panic(err)
		}
		fmt.Print(s)
	},
}

func init() {
	rootCmd.AddCommand(workerCmd)

}
