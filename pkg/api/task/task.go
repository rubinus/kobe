package task

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	"kobe/pkg/models"
	"net/http"
)

const (
	taskQueueKey = "queue"
	taskSetKey   = "task"
)

// @Summary Get Task Info
// @Tags task
// @Description Get task info
// @Param uid path string true "task_uid"
// @Produce json
// @Success 201 {object} models.Task
// @Router /tasks/{uid} [get]
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

// @Summary List Task Info
// @Description List task info
// @Tags task
// @Produce json
// @Success 200 {object} models.Task
// @Router /tasks/ [get]
func List(ctx *gin.Context) {
	r := ctx.MustGet("redis").(*redis.Client)
	ts, err := r.HGetAll(taskSetKey).Result()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	tasks := []models.Task{}
	for k, _ := range ts {
		var task models.Task
		_ = json.Unmarshal([]byte(ts[k]), &task)
		tasks = append(tasks, task)
	}
	ctx.JSON(http.StatusOK, tasks)
}
