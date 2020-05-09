package inventory

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"kobe/pkg/config"
	"kobe/pkg/models"
	"kobe/pkg/redis"
	"os"
	"testing"
)

func setInventory() string {
	config.InitConfig()
	redis.InitRedis()
	i := models.Inventory{
		Hosts: []models.Host{
			{
				Name:     "master1",
				Ip:       "172.16.10.63",
				User:     "root",
				Password: "Calong@2015",
				Port:     22,
				Vars:     map[string]interface{}{},
			},
		},
		Groups: []models.Group{
			{
				Name:     "master",
				Hosts:    []string{"master1"},
				Children: make([]string, 0),
				Vars:     make(map[string]interface{}),
			},
		},
	}
	id := uuid.NewV4().String()
	_, err := redis.Redis.Set(id, i, -1).Result()
	if err != nil {
		panic(err)
	}
	return id
}

func TestListHandler(t *testing.T) {
	config.InitConfig()
	redis.InitRedis()
	id := setInventory()
	_ = os.Setenv("inventoryId", id)
	m, err := ListHandler()
	if err != nil {
		t.Error(err)
	}
	fmt.Print(m)
}


