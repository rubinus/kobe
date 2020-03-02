package task

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	uuid "github.com/satori/go.uuid"
	"kobe/pkg/models"
	"net/http"
	"time"
)

const (
	taskQueueKey = "queue"
)

// @ params args

func Create(ctx *gin.Context) {
	var task models.Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}
	r := ctx.MustGet("redis").(*redis.Client)
	task.Uid = uuid.NewV4().String()
	task.CreatedTime = time.Now()
	task.WebSocket = fmt.Sprintf("ws//%s", task.Uid)
	task.State = models.TaskStatePending
	r.HSet(task.Uid, task)
	r.LPush(taskQueueKey, task.Uid)
	ctx.JSON(http.StatusCreated, task)
}

func Get(ctx *gin.Context) {
	r := ctx.MustGet("redis").(*redis.Client)
	uid := ctx.Param("uid")
	m := r.HGetAll(uid)
	ctx.JSON(http.StatusCreated, m)
}

func List(ctx *gin.Context) {
	r := ctx.MustGet("redis").(*redis.Client)
	ts := r.LRange(taskQueueKey, 0, -1)
	ctx.JSON(http.StatusCreated, ts)
}
