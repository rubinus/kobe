package api

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"kobe/pkg/ansible"
	"kobe/pkg/redis"
	"kobe/pkg/models"
	"net/http"
	"time"
)

type RunPlaybookRequest struct {
	Inventory models.Inventory `json:"inventory"`
	PlayBook  string           `json:"playbook"`
}

type RunPlaybookResponse struct {
	TaskId string `json:"task_id"`
}

func RunPlaybook(ctx *gin.Context) {
	var req RunPlaybookRequest
	projectName := ctx.GetString("project")
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	inventoryId, err := req.Inventory.SaveToCache()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
	}
	playbook := models.Playbook(req.PlayBook)
	project := models.Project{
		Name:      projectName,
		Playbooks: nil,
	}
	r := ansible.PlaybookRunner{Project: project}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
	}
	taskId := uuid.NewV4().String()
	go func() {
		result := &models.Result{
			StartTime: time.Time{},
			EndTime:   time.Time{},
			Message:   "",
			Success:   false,
			Content:   nil,
		}
		r.Run(inventoryId, playbook, result)
		_, _ = redis.Redis.Set(taskId, result, -1).Result()
	}()

	ctx.JSON(http.StatusOK, RunPlaybookResponse{TaskId: taskId})
}
