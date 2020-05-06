package result

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"kobe/pkg/connections"
	"kobe/pkg/models"
	"net/http"
)

const (
	resultKey = "result"
)

func Get(ctx *gin.Context) {
	uid := ctx.Param("id")
	t, err := connections.Redis.HGet(resultKey, uid).Result()
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
