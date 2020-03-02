package models

import (
	"kobe/pkg/ansible"
	"time"
)

type Host struct {
	Name     string                 `json:"name"`
	Ip       string                 `json:"ip" `
	Username string                 `json:"username"`
	Password string                 `json:"password"`
	Port     int                    `json:"port"`
	Vars     map[string]interface{} `json:"vars"`
}

type Group struct {
	Name     string                 `json:"name"`
	Children []string               `json:"children" `
	Hosts    []string               `json:"hosts"`
	Vars     map[string]interface{} `json:"vars"`
}

type Inventory struct {
	Name        string    `json:"name" `
	Groups      []Group   `json:"groups"`
	Hosts       []Host    `json:"hosts"`
	CreatedTime time.Time `json:"created_time"`
}

func (i *Inventory) Base() *ansible.BaseInventory {
	bp := &ansible.BaseInventory{
		Hosts:  nil,
		Groups: nil,
	}
	baseHosts := []ansible.BaseHost{}
	for _, h := range i.Hosts {
		vars := map[string]interface{}{
			"ansible_ssh_host": h.Ip,
			"ansible_ssh_user": h.Username,
			"ansible_ssh_pass": h.Password,
			"ansible_ssh_port": h.Port,
		}
		baseHost := ansible.BaseHost{
			Hostname: h.Name,
			Vars:     vars,
		}
		baseHosts = append(baseHosts, baseHost)
	}
	baseGroups := []ansible.BaseGroup{}
	for _, g := range i.Groups {
		hostsMap := map[string]interface{}{}
		for _, host := range g.Hosts {
			hostsMap[host] = map[string]interface{}{}
		}
		childrenMap := map[string]interface{}{}
		baseGroup := ansible.BaseGroup{
			Name:     g.Name,
			Vars:     g.Vars,
			Hosts:    hostsMap,
			Children: childrenMap,
		}
		baseGroups = append(baseGroups, baseGroup)
	}
	bp.Hosts = baseHosts
	bp.Groups = baseGroups
	return bp
}
