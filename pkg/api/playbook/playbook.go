package playbook

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"kobe/pkg/logger"
	"kobe/pkg/models"
	"net/http"
	"os"
	"path"
)

var log = logger.Logger

// @Summary List playbooks under dir
// @Description List all playbooks under dir
// @Param dir path string true "dir"
// @Produce  json
// @Tags playbook
// @Success 200 {array} models.Playbook
// @Router /playbooks/{dir} [get]
func ListByDir(ctx *gin.Context) {
	dir := ctx.Param("dir")
	ps, err := lookUpDir(dir)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, ps)
}

// @Summary List playbooks
// @Description List all playbooks
// @Produce  json
// @Tags playbook
// @Success 200 {array} models.Playbook
// @Router /playbooks/ [get]
func List(ctx *gin.Context) {
	ps, err := lookAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, ps)
}

func lookAll() ([]models.Playbook, error) {
	pwd, _ := os.Getwd()
	playbookPath := path.Join(pwd, "data", "playbooks")
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
	for _, dir := range rd {
		if dir.IsDir() {
			pss, err := lookUpDir(dir.Name())
			if err != nil {
				return nil, err
			}
			ps = append(ps, pss...)
		}

	}
	return ps, nil
}

func lookUpDir(dir string) ([]models.Playbook, error) {
	pwd, _ := os.Getwd()
	dirPath := path.Join(pwd, "data", "playbooks", dir)
	if _, err := os.Stat(dirPath); err != nil {
		if os.IsNotExist(err) {
			return nil, errors.New(fmt.Sprintf("can not find dir : %s", dirPath))
		}
	}
	rd, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.Errorf("can not read  dir %s reason %s", dirPath, err.Error())
		return nil, err
	}
	ps := make([]models.Playbook, 0)
	for _, rd := range rd {
		if !rd.IsDir() {
			ps = append(ps, models.Playbook{
				Name: rd.Name(),
				Dir:  dir,
			})
		}
	}
	log.Debugf("find %d in dir: *s", len(rd), dirPath)
	return ps, nil

}
