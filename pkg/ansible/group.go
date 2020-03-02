package ansible

type BaseGroup struct {
	Name     string
	Vars     map[string]interface{}
	Hosts    map[string]interface{}
	Children map[string]interface{}
}

func (g *BaseGroup) Data() map[string]map[string]interface{} {
	data := map[string]map[string]interface{}{}
	data[g.Name] = map[string]interface{}{
		"vars":     g.Vars,
		"hosts":    g.Hosts,
		"children": g.Children,
	}
	return data
}
