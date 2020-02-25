package inventory

type Children interface{

}

type Group struct {
	Name     string                 `json:"name"`
	Children []Children             `json:"children"`
	Vars     map[string]interface{} `json:"vars"`
}
