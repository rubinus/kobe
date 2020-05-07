package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"kobe/pkg/config"
	"kobe/pkg/redis"
	"kobe/pkg/routers"
)

func main() {
	redis.InitRedis()
	config.InitConfig()
	app := gin.Default()
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routers.InitRouter(app)
	bind := viper.GetString("server.bind")
	err := app.Run(bind)
	if err != nil {
		panic(err)
	}
}
