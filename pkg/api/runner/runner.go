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

// @Summary ImRunAdhoc
// @Tags runner
// @Description Run Adhoc Task with Inventory Object
// @Accept  json
// @Param data body models.ImRunAdhocRequest  true "request"
// @Produce json
// @Success 201 {object} models.Task
// @Router /runner/im/adhoc/ [post]
func ImRunAdhoc(ctx *gin.Context) {
	var tr models.ImRunAdhocRequest
	if err := ctx.ShouldBindJSON(&tr); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	r := ctx.MustGet("redis").(*redis.Client)
	var task models.Task
	task.Uid = uuid.NewV4().String()
	task.Args = map[string]string{
		"pattern": tr.Pattern,
		"module":  tr.Module,
		"arg":     tr.Arg,
	}
	task.CreatedTime = time.Now()
	task.Type = "adhoc"
	task.Inventory = tr.Inventory
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

// @Summary ImRunPlaybook
// @Tags runner
// @Description Create Run Playbook Task with Inventory Object
// @Accept  json
// @Param data body models.RunPlaybookRequest  true "request"
// @Produce json
// @Success 201 {object} models.Task
// @Router /runner/im/playbook/ [post]
func ImRunPlaybook(ctx *gin.Context) {
	var tr models.ImRunPlaybookRequest
	if err := ctx.ShouldBindJSON(&tr); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	r := ctx.MustGet("redis").(*redis.Client)
	var task models.Task
	task.Uid = uuid.NewV4().String()
	task.Args = map[string]string{
		"playbook": tr.Playbook,
		"dir":      tr.Dir,
	}
	task.Type = "playbook"
	task.CreatedTime = time.Now()
	task.State = models.TaskStatePending
	if _, err := r.HSet(taskSetKey, task.Uid, task).Result(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	if _, err := r.LPush(taskQueueKey, task.Uid).Result(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, task)
}
