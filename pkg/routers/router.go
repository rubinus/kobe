package routers

import (
	"github.com/gin-gonic/gin"
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
		t := v1.Group("tasks")
		{
			t.GET("/", task.List)
			t.GET("/:uid", task.Get)
		}
		ru := v1.Group("/runner")
		{

		}
		im := ru.Group("/im")
		{
			im.POST("/playbook/", runner.ImRunPlaybook)
			im.POST("/adhoc/", runner.ImRunAdhoc)
		}
		r := v1.Group("result")
		{
			r.GET("/:uid", result.Get)
		}
	}
}
