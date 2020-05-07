package models

import (
	"encoding/json"
	uuid "github.com/satori/go.uuid"
	"kobe/pkg/redis"
)

type Host struct {
	Ip         string                 `json:"ip"`
	Name       string                 `json:"name"`
	Port       int                    `json:"port"`
	User       string                 `json:"user"`
	Password   string                 `json:"password"`
	Connection string                 `json:"connection"`
	Vars       map[string]interface{} `json:"vars"`
}

type Group struct {
	Name     string                 `json:"name"`
	Hosts    []string               `json:"hosts"`
	Children []string               `json:"children"`
	Vars     map[string]interface{} `json:"vars"`
}

type Inventory struct {
	Hosts  []Host  `json:"hosts"`
	Groups []Group `json:"groups"`
}

func (i Inventory) MarshalBinary() (data []byte, err error) {
	return json.Marshal(i)
}

func (i Inventory) SaveToCache() (string, error) {
	id := uuid.NewV4().String()
	_, err := redis.Redis.Set(id, i, -1).Result()
	if err != nil {
		return "", err
	}
	return id, nil
}
