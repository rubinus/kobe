package inventory

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"kobe/pkg/models"
	"net/http"
)

const modelName = "inventory"

func Create(ctx *gin.Context) {
	db := ctx.MustGet("db").(*mgo.Database)
	var i models.Inventory
	if err := ctx.ShouldBind(&i); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c := db.C(modelName)
	index := mgo.Index{
		Key:    []string{"name"},
		Unique: true,
	}
	_ = c.EnsureIndex(index)
	if err := c.Insert(i); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, i)
}

func Update(ctx *gin.Context) {
	db := ctx.MustGet("db").(*mgo.Database)
	var i models.Inventory
	name := ctx.Param("name")
	if err := ctx.ShouldBind(&i); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c := db.C(modelName)
	i.Name = name
	if err := c.Update(bson.M{"name": i.Name}, i); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusAccepted, i)
}
func List(ctx *gin.Context) {
	db := ctx.MustGet("db").(*mgo.Database)
	c := db.C(modelName)
	i := []models.Inventory{}
	if err := c.Find(nil).Sort("-created_time").All(&i); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	ctx.JSON(http.StatusOK, i)
}

func Get(ctx *gin.Context) {
	db := ctx.MustGet("db").(*mgo.Database)
	c := db.C(modelName)
	i := models.Inventory{}
	name := ctx.Param("name")
	query := bson.M{"name": name}
	if err := c.Find(query).One(&i); err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, i)
}

func Delete(ctx *gin.Context) {
	db := ctx.MustGet("db").(*mgo.Database)
	name := ctx.Param("name")
	query := bson.M{"name": name}
	err := db.C(modelName).Remove(query)
	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"name": name})
}
