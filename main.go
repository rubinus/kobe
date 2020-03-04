package main

import "kobe/cmd"
import _ "kobe/docs"

// @title Kobe Restful API
// @version 0.0.1
// @description  This is RestAPI Client for ansible
// @BasePath /api/v1/
func main() {
	cmd.Execute()
}
