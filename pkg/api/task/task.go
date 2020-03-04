package task

import (
	"encoding/json"
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

func Create(ctx *gin.Context) {
	var task models.Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}
	r := ctx.MustGet("redis").(*redis.Client)
	task.Uid = uuid.NewV4().String()
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

func Get(ctx *gin.Context) {
	r := ctx.MustGet("redis").(*redis.Client)
	uid := ctx.Param("uid")
	t, err := r.HGet(taskSetKey, uid).Result()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	var task models.Task
	if err := json.Unmarshal([]byte(t), &task); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, task)
}

func List(ctx *gin.Context) {
	r := ctx.MustGet("redis").(*redis.Client)
	ts, err := r.HGetAll(taskSetKey).Result()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, ts)
}
