package routers

import (
	"github.com/gin-gonic/gin"
	"kobe/pkg/api/inventory"
	"kobe/pkg/api/playbook"
	"kobe/pkg/api/result"
	"kobe/pkg/api/runner"
	"kobe/pkg/api/task"
)

func InitRouter(g *gin.Engine) {
	v1 := g.Group("/api/v1")
	{
		p := v1.Group("/playbooks")
		{
			p.GET("/", playbook.List)
			p.GET("/:dir", playbook.ListByDir)
		}
		i := v1.Group("/inventory")
		{
			i.GET("/", inventory.List)
			i.POST("/", inventory.Create)
			i.PUT("/:name", inventory.Update)
			i.GET("/:name", inventory.Get)
			i.DELETE("/:name", inventory.Delete)
		}
		t := v1.Group("tasks")
		{
			t.GET("/", task.List)
			t.GET("/:uid", task.Get)
			t.POST("/playbook/", runner.RunPlaybook)
			t.POST("/adhoc/", runner.RunAdhoc)
		}
		r := v1.Group("result")
		{
			r.GET("/:uid", result.Get)
		}
	}
}
