package result

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	"kobe/pkg/models"
	"net/http"
)

const (
	resultKey = "result"
)

// @Summary Get Task Result
// @Description Get task result by task id when task finished
// @Param uid path string true "task_uid"
// @Produce json
// @Success 201 {object} models.Result
// @Router /result/{uid} [get]
func Get(ctx *gin.Context) {
	r := ctx.MustGet("redis").(*redis.Client)
	uid := ctx.Param("uid")
	t, err := r.HGet(resultKey, uid).Result()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	var result models.Result
	if err := json.Unmarshal([]byte(t), &result); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, result)
}
