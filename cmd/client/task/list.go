package task

import (
	"fmt"
	"github.com/spf13/cobra"
	"kobe/pkg/client"
	"log"
)

var taskListCmd = &cobra.Command{
	Use: "list",
	Run: func(cmd *cobra.Command, args []string) {
		c := client.NewKobeClient("127.0.0.1", 8080)
		rs, err := c.ListResult()
		if err != nil {
			log.Fatal(err)
		}
		for _, r := range rs {
			out := fmt.Sprintf("%s  %s   %s   %t  %t",
				r.Id, r.StartTime, r.EndTime, r.Finished, r.Success)
			fmt.Println(out)
		}
	},
}

func init() {
}
