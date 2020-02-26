package server

import "github.com/kataras/iris"
import _ "kobe/pkg/broker"

var App = iris.New()

func RunServer() error {
	App.Logger().SetLevel("debug")
	return App.Run(iris.Addr(":8080"))
}
