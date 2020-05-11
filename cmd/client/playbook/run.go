package playbook

import (
	"fmt"
	"github.com/spf13/cobra"
	"kobe/api"
	"kobe/client"
	"log"
	"os"
)

var playbookRunCmd = &cobra.Command{
	Use: "run",
	Run: func(cmd *cobra.Command, args []string) {
		c := client.NewKobeClient("127.0.0.1", 8080)
		project, err := cmd.Flags().GetString("project")
		if err != nil {
			log.Fatal(err)
		}
		if len(args) < 1 {
			log.Fatal("invalid playbook name")
		}
		playbook := args[0]
		var result api.Result
		i := api.Inventory{
			Hosts: []*api.Host{
				{
					Name:     "test",
					Ip:       "172.16.10.63",
					Port:     22,
					User:     "root",
					Password: "Calong@2015",
					Vars:     map[string]string{},
				},
			},
			Groups: []*api.Group{
				{
					Name:     "master",
					Children: []string{},
					Hosts:    []string{"test"},
				},
			},
		}
		err = c.RunPlaybook(project, playbook, i, os.Stdout, &result)
		fmt.Print(result.Success)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	playbookRunCmd.Flags().StringP("project", "p", "", "")
}
