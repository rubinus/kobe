package runner

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	uuid "github.com/satori/go.uuid"
	"kobe/pkg/models"
	"net/http"
	"time"
)

const (
	taskQueueKey = "queue"
	taskSetKey   = "task"
)

// @Summary RunAdhoc
// @Description Create Run Adhoc Task
// @Accept  json
// @Param data body models.RunAdhocRequest  true "create adhoc task"
// @Produce json
// @Success 201 {object} models.Task
// @Router /tasks/adhoc/ [post]
func RunAdhoc(ctx *gin.Context) {
	var tr models.RunAdhocRequest
	if err := ctx.ShouldBindJSON(&tr); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}
	r := ctx.MustGet("redis").(*redis.Client)
	var task models.Task
	task.Uid = uuid.NewV4().String()
	task.Args = map[string]string{
		"inventory": tr.Inventory,
		"pattern":   tr.Pattern,
		"module":    tr.Module,
		"arg":       tr.Arg,
	}
	task.CreatedTime = time.Now()
	task.Type = "adhoc"
	task.State = models.TaskStatePending
	if _, err := r.HSet(taskSetKey, task.Uid, task).Result(); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if _, err := r.LPush(taskQueueKey, task.Uid).Result(); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, task)
}

// @Summary RunPlaybook
// @Description Create Run Playbook Task
// @Accept  json
// @Param data body models.RunPlaybookRequest  true "create playbook task"
// @Produce json
// @Success 201 {object} models.Task
// @Router /tasks/playbook/ [post]
func RunPlaybook(ctx *gin.Context) {
	var tr models.RunPlaybookRequest
	if err := ctx.ShouldBindJSON(&tr); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}
	r := ctx.MustGet("redis").(*redis.Client)
	var task models.Task
	task.Uid = uuid.NewV4().String()
	task.Args = map[string]string{
		"inventory": tr.Inventory,
		"playbook":  tr.Playbook,
	}
	task.Type = "playbook"
	task.CreatedTime = time.Now()
	task.State = models.TaskStatePending
	if _, err := r.HSet(taskSetKey, task.Uid, task).Result(); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if _, err := r.LPush(taskQueueKey, task.Uid).Result(); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, task)
}
