package ansible

type BaseInventory struct {
	Hosts  []BaseHost
	Groups []BaseGroup
}

func (bi BaseInventory) Data() map[string]map[string]interface{} {
	data := map[string]map[string]interface{}{}
	allGroup := BaseGroup{
		Name:  "all",
		Hosts: map[string]interface{}{},
	}
	localhost := BaseHost{
		Hostname: "localhost",
		Vars: map[string]interface{}{
			"ansible_connection": "local",
			"ansible_ssh_port":   22,
			"ansible_ssh_user":   "root",
		},
	}
	hosts := append(bi.Hosts, localhost)
	hostsMap := map[string]interface{}{}
	for _, host := range hosts {
		for k, v := range host.Data() {
			hostsMap[k] = v
		}
	}
	allGroup.Hosts = hostsMap
	groups := append(bi.Groups, allGroup)
	for _, group := range groups {
		for k, v := range group.Data() {
			data[k] = v
		}
	}
	return data
}
