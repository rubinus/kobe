package playbook

import (
	"fmt"
	"github.com/spf13/cobra"
	"kobe/pkg/client"
	"log"
)

var playbookListCmd = &cobra.Command{
	Use: "list",
	Run: func(cmd *cobra.Command, args []string) {
		c := client.NewKobeClient("127.0.0.1", 8080)
		ps, err := c.ListProject()
		if err != nil {
			log.Fatal(err.Error())
		}
		for _, p := range ps {
			fmt.Println(p.Name)
			for _, pb := range p.Playbooks {
				str := fmt.Sprintf("-- %s", pb)
				fmt.Println(str)
			}
		}
	},
}

func init() {
}
