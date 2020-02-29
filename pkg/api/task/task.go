package task

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"kobe/pkg/models"
	"net/http"
)

const modelName = "task"

func Create(ctx *gin.Context) {
	db := ctx.MustGet("db").(*mgo.Database)
	var t models.Task
	if err := ctx.ShouldBind(&t); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c := db.C(modelName)
	index := mgo.Index{
		Key:    []string{"uid"},
		Unique: true,
	}
	_ = c.EnsureIndex(index)
	t.State = models.TaskStateScheduling
	t.Success = false
	t.Scheduled = false

	if err := c.Insert(t); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, t)
}

func Get(ctx *gin.Context) {
	db := ctx.MustGet("db").(*mgo.Database)
	c := db.C(modelName)
	t := models.Task{}
	uid := ctx.Param("uid")
	query := bson.M{"uid": uid}
	if err := c.Find(query).One(&t); err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, t)
}

func List(ctx *gin.Context) {
	db := ctx.MustGet("db").(*mgo.Database)
	ps := []models.Task{}
	err := db.C(modelName).Find(nil).Sort("-created_data").All(&ps)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, ps)
}
