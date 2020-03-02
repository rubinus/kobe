package middlewares

import (
	"github.com/gin-gonic/gin"
	"kobe/pkg/connections"
	"kobe/pkg/logger"
)

var log = logger.Logger

func Connect(c *gin.Context) {
	s := connections.Redis
	defer s.Close()
	c.Set("redis", s)
	c.Next()
}
