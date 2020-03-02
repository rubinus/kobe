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

func lookUp() ([]models.PlaybookSet, error) {
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
	pss := make([]models.PlaybookSet, 0)
	for _, r := range rd {
		if r.IsDir() {
			ps := models.PlaybookSet{
				Name: r.Name(),
				Path: path.Join(playbookPath, r.Name()),
			}
			setPath := path.Join(playbookPath, r.Name())
			sd, err := ioutil.ReadDir(setPath)
			if err != nil {
				log.Errorf("can not read playbookSet dir %s reason %s", setPath, err.Error())
				continue
			}
			pbs := make([]models.Playbook, 0)
			for _, s := range sd {
				if !s.IsDir() {
					p := models.Playbook{Name: s.Name()}
					pbs = append(pbs, p)
				}
			}
			pss = append(pss, ps)
		}
	}
	return pss, nil
}
