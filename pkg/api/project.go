package api

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"kobe/pkg/constant"
	"kobe/pkg/models"
	"kobe/pkg/util"
	"net/http"
	"path"
	"strings"
)

type CreateProjectRequest struct {
	Name   string
	Source string
}

type CreateProjectResponse struct {
	Name string
}

func CreateProject(ctx *gin.Context) {
	var req CreateProjectRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	destPath := path.Join(constant.ProjectDir, req.Name)
	if err := util.CloneRepository(req.Source, destPath); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, CreateProjectResponse{
		Name: req.Name,
	})
}

func ListProject(ctx *gin.Context) {
	ps, err := lookUpStorage()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, ps)
}

func lookUpStorage() ([]models.Project, error) {
	var results []models.Project
	rd, err := ioutil.ReadDir(constant.ProjectDir)
	if err != nil {
		return nil, err
	}
	for _, d := range rd {
		if d.IsDir() {
			playbooks, err := searchPlays(path.Join(constant.ProjectDir, d.Name()))
			if err != nil {
				return nil, err
			}
			results = append(results, models.Project{
				Name:      d.Name(),
				Playbooks: playbooks,
			})
		}
	}
	return results, nil
}

func searchPlays(parent string) ([]models.Playbook, error) {
	var results []models.Playbook
	rd, err := ioutil.ReadDir(parent)
	if err != nil {
		return nil, err
	}
	for _, f := range rd {
		if !f.IsDir() && strings.HasSuffix(f.Name(), ".yml") {
			results = append(results, models.Playbook(f.Name()))
		}
	}
	return results, nil
}
