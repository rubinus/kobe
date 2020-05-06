package routers

import (
	"github.com/gin-gonic/gin"
	"kobe/pkg/api"
	"kobe/pkg/api/result"
	"kobe/pkg/api/runner"
)

func InitRouter(g *gin.Engine) {
	v1 := g.Group("/api/v1")
	{
		p := v1.Group("/projects")
		{
			p.POST("/", api.CreateProject)
			p.GET("/", api.ListProject)
		}
		ru := v1.Group("/runner")
		{
			ru.POST("/", runner.RunPlaybook)
		}
		r := v1.Group("result")
		{
			r.GET("/:id", result.Get)
		}
	}
}
