package inventory

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

// @Summary Create Inventory
// @Description Create Inventory
// @Accept  json
// @Param data body models.Inventory  true "create inventory"
// @Produce json
// @Tags inventory
// @Success 201 {object} models.Inventory
// @Router /inventory/ [post]
func Create(ctx *gin.Context) {
	var i models.Inventory
	if err := ctx.ShouldBind(&i); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	e, err := exists(i.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	if e {
		msg := fmt.Sprintf("%s already exists", i.Name)
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": msg})
		return
	}
	if err := write(i.Name, i.Content); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, i)
}

// @Summary Update Inventory
// @Description Update Inventory
// @Accept  json
// @Tags inventory
// @Param data body models.Inventory  true "update inventory"
// @Param name path string true "name"
// @Produce json
// @Success 202 {object} models.Inventory
// @Router /inventory/{name} [put]
func Update(ctx *gin.Context) {
	var i models.Inventory
	if err := ctx.ShouldBind(&i); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	e, err := exists(i.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	if !e {
		msg := fmt.Sprintf("%s not exists", i.Name)
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": msg})
		return
	}
	if err := write(i.Name, i.Content); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	ctx.JSON(http.StatusAccepted, i)
}

// @Summary List Inventory
// @Description List Inventory
// @Produce json
// @Tags inventory
// @Success 200 {array} models.Inventory
// @Router /inventory/ [get]
func List(ctx *gin.Context) {
	is, err := lookUp()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, is)
}

// @Summary Get Inventory
// @Description Get Inventory
// @Produce json
// @Tags inventory
// @Param name path string true "name"
// @Success 200 {object} models.Inventory
// @Router /inventory/{name} [get]
func Get(ctx *gin.Context) {
	name := ctx.Param("name")
	i, err := GetModel(name)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, i)
}

// @Summary Delete Inventory
// @Description Delete Inventory
// @Produce json
// @Tags inventory
// @Param name path string true "name"
// @Success 200 {string} string "name"
// @Router /inventory/{name} [delete]
func Delete(ctx *gin.Context) {
	name := ctx.Param("name")
	e, err := exists(name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	if !e {
		msg := fmt.Sprintf("can not find %s", name)
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": gin.H{"msg": msg}})
		return
	}
	if err := remove(""); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"name": name})
}

func lookUp() ([]models.Inventory, error) {
	pwd, _ := os.Getwd()
	inventoryPath := path.Join(pwd, "data", "inventory")
	log.Debugf("search inventory in path: %s", inventoryPath)
	if err := os.MkdirAll(inventoryPath, 0755); err != nil {
		log.Errorf("can not make inventory dir %s reason %s", inventoryPath, err.Error())
		return nil, err
	}
	rd, err := ioutil.ReadDir(inventoryPath)
	if err != nil {
		log.Errorf("can not read playbook dir %s reason %s", inventoryPath, err.Error())
		return nil, err
	}
	is := make([]models.Inventory, 0)
	for _, r := range rd {
		if !r.IsDir() {
			content, err := read(r.Name())
			if err != nil {
				continue
			}
			i := models.Inventory{
				Name:    r.Name(),
				Content: content,
			}
			is = append(is, i)
		}
	}
	return is, nil
}

func read(name string) (string, error) {
	pwd, _ := os.Getwd()
	inventoryPath := path.Join(pwd, "data", "inventory")
	bs, err := ioutil.ReadFile(path.Join(inventoryPath, name))
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

func write(name string, content string) error {
	pwd, _ := os.Getwd()
	inventoryPath := path.Join(pwd, "data", "inventory")
	err := ioutil.WriteFile(path.Join(inventoryPath, name), []byte(content), 0755)
	if err != nil {
		return err
	}
	return nil
}

func exists(name string) (bool, error) {
	pwd, _ := os.Getwd()
	inventoryPath := path.Join(pwd, "data", "inventory")
	rd, err := ioutil.ReadDir(inventoryPath)
	if err != nil {
		return false, err
	}
	for _, f := range rd {
		if f.Name() == name {
			return true, nil
		}
	}
	return false, nil
}

func GetModel(name string) (*models.Inventory, error) {
	pwd, _ := os.Getwd()
	inventoryPath := path.Join(pwd, "data", "inventory")
	rd, err := ioutil.ReadDir(inventoryPath)
	if err != nil {
		return nil, err
	}
	for _, r := range rd {
		if r.Name() == name {
			bytes, err := ioutil.ReadFile(path.Join(inventoryPath, name))
			if err != nil {
				return nil, err
			}
			i := &models.Inventory{
				Name:    name,
				Content: string(bytes),
			}
			return i, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("can not find %s", name))
}

func remove(name string) error {
	pwd, _ := os.Getwd()
	inventoryPath := path.Join(pwd, "data", "inventory")
	rd, err := ioutil.ReadDir(inventoryPath)
	if err != nil {
		return err
	}
	for _, r := range rd {
		if r.Name() == name {
			return os.Remove(path.Join(inventoryPath, r.Name()))
		}
	}
	return nil
}
