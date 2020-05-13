package project

import (
	"fmt"
	"github.com/spf13/cobra"
	"kobe/pkg/client"
	"log"
)

var projectCreateCmd = &cobra.Command{
	Use: "create",
	Run: func(cmd *cobra.Command, args []string) {
		c := client.NewKobeClient("127.0.0.1", 8080)
		if len(args) < 0 {
			log.Fatal("invalid project source")
		}
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			log.Fatal(err)
		}
		source := args[0]
		p, err := c.CreateProject(name, source)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(fmt.Sprintf("project %s created", p.Name))

	},
}

func init() {
	projectCreateCmd.Flags().String("name", "", "")
}
