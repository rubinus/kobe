package models

import (
	"encoding/json"
	"fmt"
)

type Playbook struct {
	Name string `json:"name"`
	Dir  string `json:"dir"`
}

func (p *Playbook) MarshalBinary() (data []byte, err error) {
	return json.Marshal(p)
}

func (p Playbook) String() string {
	return fmt.Sprintf("%s@%s", p.Dir, p.Name)
}
