package ansible

type BaseGroup struct {
    Name     string
    Vars     map[string]interface{}
    Hosts    []BaseHost
    Children []BaseGroup
}

func (bg BaseGroup) Data() map[string]interface{} {
    groupData := make(map[string]interface{})
    groupData["hosts"] = make([]string, 0)
    groupData["children"] = make([]string, 0)
    for _, host := range bg.Hosts {
        groupData["hosts"] = host.Data()
    }
    for _, children := range bg.Children {
        groupData["children"] = children.Data()
    }
    groupData["vars"] = bg.Vars
    return groupData
}