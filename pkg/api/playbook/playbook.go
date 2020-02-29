package playbook

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"kobe/pkg/logger"
	"kobe/pkg/models"
	"net/http"
	"os"
	"path"
	"time"
)

var log = logger.Logger

const modelName = "playbook"

func List(ctx *gin.Context) {
	db := ctx.MustGet("db").(*mgo.Database)
	c := db.C(modelName)
	index := mgo.Index{
		Key:    []string{"name"},
		Unique: true,
	}
	_ = c.EnsureIndex(index)
	if err := lookUp(db); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ps := []models.Playbook{}
	err := c.Find(nil).Sort("-created_time").All(&ps)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, ps)
}

func lookUp(db *mgo.Database) error {
	c := db.C(modelName)
	pwd, _ := os.Getwd()
	playbookPath := path.Join(pwd, "data", "playbooks")
	log.Debugf("search playbook in path: %s", playbookPath)
	if err := os.MkdirAll(playbookPath, 0755); err != nil {
		return err
	}
	rd, err := ioutil.ReadDir(playbookPath)
	if err != nil {
		return err
	}
	for _, r := range rd {
		if r.IsDir() {
			p := models.Playbook{
				Name:        r.Name(),
				Path:        path.Join(playbookPath, r.Name()),
				CreatedTime: time.Now(),
			}
			if _, err := c.Upsert(bson.M{"name": p.Name}, p); err != nil {
				return err
			}
			log.Debugf("discover playbook %s in path %s", p.Name, p.Path)
		}
	}
	return nil
}
