package server

import (
	"github.com/kataras/iris"
)

func RunServer() error {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Get("/ping", func(ctx iris.Context) {
		_, _ = ctx.WriteString("pong")
	})
	if err := app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed)); err != nil {
		return err
	}
	return nil
}
