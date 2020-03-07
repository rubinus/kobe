package server

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"kobe/pkg/connections"
	"kobe/pkg/middlewares"
	"kobe/pkg/routers"
)

func Run() error {
	connections.ConnectRedis()
	app := gin.Default()
	app.Use(middlewares.SetRedis)
	app.Delims("{[", "]}")
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routers.InitRouter(app)
	bind := viper.GetString("server.bind")
	return app.Run(bind)
}
