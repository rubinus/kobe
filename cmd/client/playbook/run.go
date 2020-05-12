package playbook

import (
	"fmt"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"io/ioutil"
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
		inventoryPath, err := cmd.Flags().GetString("inventory")
		if err != nil {
			log.Fatal(err)
		}
		content, err := ioutil.ReadFile(inventoryPath)
		if err != nil {
			log.Fatal(err)
		}
		var inventory api.Inventory
		err = yaml.Unmarshal(content, &inventory)
		if err != nil {
			log.Fatal(err)
		}
		if len(args) < 1 {
			log.Fatal("invalid playbook name")
		}
		playbook := args[0]
		result, err := c.RunPlaybook(project, playbook, inventory)
		if err != nil {
			log.Fatal(err)
		}
		backend, err := cmd.Flags().GetBool("b")
		if err != nil {
			log.Fatal(err)
		}
		if backend {
			fmt.Println(result.Id)
		} else {
			err := c.WatchRunPlaybook(result.Id, os.Stdout)
			if err != nil {
				log.Fatal(err)
			}
		}

	},
}

func init() {
	playbookRunCmd.Flags().StringP("project", "p", "", "")
	playbookRunCmd.Flags().BoolP("b", "b", false, "")
	playbookRunCmd.Flags().StringP("inventory", "i", "", "")
}
