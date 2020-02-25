package ansible

import (
    "encoding/json"
    "log"
)

type BaseInventory struct {
    Hosts  []BaseHost
    Groups []BaseGroup
}

func NewBaseInventory(hosts []BaseHost, groups []BaseGroup) *BaseInventory {
    i := BaseInventory{Hosts: hosts, Groups: groups}
    return &i
}

func (bi BaseInventory) Data() map[string]interface{} {
    localhost := BaseHost{
        Hostname: "localhost",
        Vars: map[string]interface{}{
            "ansible_ssh_host":   "172.0.0.1",
            "ansible_connection": "local",
        },
    }
    hosts := append(bi.Hosts, localhost)
    allGroup := BaseGroup{
        Name:     "all",
        Vars:     map[string]interface{}{},
        Hosts:    hosts,
        Children: []BaseGroup{},
    }
    groups := append(bi.Groups, allGroup)
    inventoryData := make(map[string]interface{})
    for _, group := range groups {
        inventoryData[group.Name] = group.Data()
    }
    return inventoryData
}

func (bi BaseInventory) String() string {
    bytes, err := json.Marshal(bi.Data())
    if err != nil {
        log.Printf("parse json error: %s", bytes)
    }
    return string(bytes)
}
