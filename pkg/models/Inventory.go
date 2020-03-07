package models

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
	Hosts    []Host                 `json:"hosts"`
	Children []string               `json:"children"`
	Vars     map[string]interface{} `json:"vars"`
}

type Inventory struct {
	Hosts  []Host  `json:"hosts"`
	Groups []Group `json:"groups"`
}

func (i *Inventory) Data() map[string]interface{} {
	allGroup := Group{
		Name: "all",
		Hosts: append(i.Hosts, Host{
			Ip:         "127.0.0.1",
			Name:       "localhost",
			Port:       22,
			Connection: "local",
			Vars:       map[string]interface{}{},
		}),
		Vars: map[string]interface{}{},
	}
	groups := append(i.Groups, allGroup)
	groupMap := map[string]interface{}{}
	for _, group := range groups {
		gm := map[string]interface{}{}
		gm["hosts"] = map[string]interface{}{}
		for _, host := range group.Hosts {
			hostMap := map[string]interface{}{}
			if host.Connection != "" {
				hostMap["ansible_connection"] = host.Connection
			}
			hostMap["ansible_ssh_host"] = host.Ip
			hostMap["ansible_ssh_pass"] = host.Password
			hostMap["ansible_port"] = host.Port
			hostMap["ansible_ssh_user"] = host.User
			hostMap["vars"] = host.Vars
			gm["hosts"].(map[string]interface{})[host.Name] = hostMap
		}
		gm["children"] = map[string]interface{}{}
		for _, c := range group.Children {
			cm := map[string]interface{}{}
			cm[c] = map[string]interface{}{}
			gm["children"].(map[string]interface{})[c] = cm
		}
		gm["vars"] = group.Vars
		groupMap[group.Name] = gm
	}
	return groupMap
}
