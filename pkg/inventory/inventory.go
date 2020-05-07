package inventory

import (
	"encoding/json"
	"errors"
	"fmt"
	"kobe/pkg/models"
	"kobe/pkg/redis"
	"os"
)

type Result map[string]map[string]interface{}

func (r Result) String() string {
	b, err := json.Marshal(&r)
	if err != nil {
		return ""
	}
	return string(b)
}

func getInventoryFromCache(id string) (*models.Inventory, error) {
	i, err := redis.Redis.Get(id).Result()
	if err != nil {
		return nil, err
	}
	var inventory models.Inventory
	if err := json.Unmarshal([]byte(i), &inventory); err != nil {
		return nil, err
	}
	return &inventory, nil
}

func HostHandler() (Result, error) {
	id, err := getInventoryIdFromEnv()
	if err != nil {
		return nil, err
	}
	inventory, _ := getInventoryFromCache(id)
	hostMap := make(map[string]map[string]interface{})
	for _, host := range inventory.Hosts {
		vars := make(map[string]interface{})
		hostMap[host.Name] = map[string]interface{}{
			"ansible_ssh_host": host.Ip,
			"ansible_ssh_port": host.Port,
			"ansible_ssh_user": host.User,
			"ansible_ssh_pass": host.Password,
		}
		if host.Vars != nil {
			for k, v := range host.Vars {
				vars[k] = v
			}
			hostMap["vars"] = vars
		}
	}
	return hostMap, nil
}

func ListHandler() (Result, error) {
	id, err := getInventoryIdFromEnv()
	if err != nil {
		return nil, err
	}
	inventory, _ := getInventoryFromCache(id)
	allGroup := make(map[string]map[string]interface{})
	for _, group := range inventory.Groups {

		m := map[string]interface{}{
			"hosts":    group.Hosts,
		}
		if group.Children != nil {
			m["children"] = group.Children
		}
		if group.Vars != nil {
			m["vars"] = group.Vars
		}
		allGroup[group.Name] = m
	}
	return allGroup, nil
}

func getInventoryIdFromEnv() (string, error) {
	id := os.Getenv("inventoryId")
	if id == "" {
		return "", errors.New(fmt.Sprintf("invalid id: %s", id))
	}
	return id, nil
}
