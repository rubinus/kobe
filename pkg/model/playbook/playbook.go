package playbook

import (
	"kobe/pkg/ansible"
)

//var cache Cache

func init() {
	//cache = Cache{
	//	items: map[string]*Playbook{},
	//}
}

type Playbook struct {
	Id   string
	Name string
	base ansible.BasePlaybook
}

//
//func (p *Playbook) Manager() {
////Manager	manager := model.Manager{
////		cache: cache,
////		t:     reflect.TypeOf(Playbook{}),
////	}
//
//}
