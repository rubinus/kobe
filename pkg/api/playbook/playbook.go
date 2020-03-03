package playbook

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"kobe/pkg/logger"
	"kobe/pkg/models"
	"net/http"
	"os"
	"path"
)

var log = logger.Logger

func List(ctx *gin.Context) {
	ps, err := lookUp()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, ps)
}

func lookUp() ([]models.Playbook, error) {
	pwd, _ := os.Getwd()
	playbookPath := path.Join(pwd, "data", "playbooks")
	log.Debugf("search playbook in path: %s", playbookPath)
	if err := os.MkdirAll(playbookPath, 0755); err != nil {
		log.Errorf("can not make playbook dir %s reason %s", playbookPath, err.Error())
		return nil, err
	}
	rd, err := ioutil.ReadDir(playbookPath)
	if err != nil {
		log.Errorf("can not read playbook dir %s reason %s", playbookPath, err.Error())
		return nil, err
	}
	ps := make([]models.Playbook, 0)
	for _, r := range rd {
		if !r.IsDir() {
			p := models.Playbook{
				Name: r.Name(),
			}
			ps = append(ps, p)
		}
	}
	return ps, nil
}
