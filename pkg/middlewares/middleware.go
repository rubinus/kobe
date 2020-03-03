package middlewares

import (
	"github.com/gin-gonic/gin"
	"kobe/pkg/connections"
	"kobe/pkg/logger"
)

var log = logger.Logger

func SetRedis(c *gin.Context) {
	s := connections.Redis
	c.Set("redis", s)
	c.Next()
}
