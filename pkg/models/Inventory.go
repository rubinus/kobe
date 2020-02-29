package models

import "time"

type Host struct {
	Name     string                 `json:"name" bson:"name"`
	Ip       string                 `json:"ip" bson:"ip"`
	Username string                 `json:"username" bson:"username"`
	Password string                 `json:"password" bson:"password"`
	Port     int                    `json:"port" bson:"port"`
	Vars     map[string]interface{} `json:"vars" bson:"vars"`
}

type Group struct {
	Name     string                 `json:"name" bson:"name"`
	Children []string               `json:"children" bson:"children"`
	Hosts    []string               `json:"hosts" bson:"hosts"`
	Vars     map[string]interface{} `json:"vars" bson:"vars"`
}

type Inventory struct {
	Id          string    `json:"-" bson:"_id,omitempty"`
	Name        string    `json:"name" bson:"name"`
	Groups      []Group   `json:"groups" bson:"groups"`
	Hosts       []Host    `json:"hosts" bson:"hosts"`
	CreatedTime time.Time `json:"created_time" bson:"created_time"`
}
