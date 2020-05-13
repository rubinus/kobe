package project

import (
	"fmt"
	"github.com/spf13/cobra"
	"kobe/pkg/client"
	"log"
)

var projectListCmd = &cobra.Command{
	Use: "list",
	Run: func(cmd *cobra.Command, args []string) {
		c := client.NewKobeClient("127.0.0.1", 8080)
		ps, err := c.ListProject()
		if err != nil {
			log.Fatal(err)
		}
		for _, p := range ps {
			fmt.Println(p.Name)
		}
	},
}

func init() {
}
