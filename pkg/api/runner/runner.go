package runner

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"kobe/pkg/ansible"
	"kobe/pkg/connections"
	"kobe/pkg/models"
	"net/http"
	"time"
)

type RunPlaybookRequest struct {
	Inventory    models.Inventory `json:"inventory"`
	PlaybookName string           `json:"playbook_name"`
	Play         string           `json:"play"`
}

type RunPlaybookResponse struct {
	TaskId string `json:"task_id"`
}

func RunPlaybook(ctx *gin.Context) {
	var req RunPlaybookRequest
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	inventoryId, err := req.Inventory.SaveToCache()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
	}
	playbook := models.Playbook{
		Name: req.PlaybookName,
	}
	r := ansible.PlaybookRunner{Playbook: playbook}
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
		r.Run(inventoryId, req.Play, result)
		_, _ = connections.Redis.Set(taskId, result, -1).Result()
	}()

	ctx.JSON(http.StatusOK, RunPlaybookResponse{TaskId: taskId})
}
