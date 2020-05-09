package main

import (
	"kobe/cmd/client/root"
	"kobe/pkg/config"
	"os"
)

func main() {
	config.InitConfig()
	if err := root.Cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
