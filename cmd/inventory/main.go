package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"kobe/pkg/config"
	"kobe/pkg/inventory"
	"kobe/pkg/redis"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "inventory",
	Short: "A inventory provider for kobe",
	Run: func(cmd *cobra.Command, args []string) {
		list, err := cmd.Flags().GetBool("list")
		host, err := cmd.Flags().GetBool("host")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if list {
			result, err := inventory.ListHandler()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			fmt.Println(result)
			os.Exit(0)
		}
		if host {
			result, err := inventory.HostHandler()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println(result)
			os.Exit(0)
		}
	},
}

func init() {
	rootCmd.Flags().Bool("host", false, "")
	rootCmd.Flags().Bool("list", false, "")
}

func main() {
	config.InitConfig()
	redis.InitRedis()
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}

}
