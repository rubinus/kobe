package task

import (
	"fmt"
	"github.com/spf13/cobra"
	"kobe/client"
	"log"
)

var taskDescribeCmd = &cobra.Command{
	Use: "describe",
	Run: func(cmd *cobra.Command, args []string) {
		c := client.NewKobeClient("127.0.0.1", 8080)
		if len(args) < 1 {
			log.Fatal("task id missing")
		}
		taskId := args[0]
		result, err := c.GetResult(taskId)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(fmt.Sprintf("id: %s", result.Id))
		fmt.Println(fmt.Sprintf("star time: %s", result.StartTime))
		fmt.Println(fmt.Sprintf("end time: %s", result.EndTime))
		fmt.Println(fmt.Sprintf("finished: %t", result.Finished))
		fmt.Println(fmt.Sprintf("success: %t", result.Success))
		fmt.Println(fmt.Sprintf("content:"))
		fmt.Println(result.Content)
	},
}

func init() {

}
