package worker

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"kobe/pkg/models"
	"net/http"
)

const modelName = "worker"

func List(ctx *gin.Context) {
	db := ctx.MustGet("db").(*mgo.Database)
	var workers []models.Worker
	err := db.C(modelName).Find(nil).All(&workers)
	if err != nil {
		err.Error()
	}
	ctx.JSON(http.StatusOK, workers)
}
