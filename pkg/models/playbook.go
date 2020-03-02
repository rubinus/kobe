package models

import (
	"kobe/pkg/ansible"
	"time"
)

type Playbook struct {
	Id          string    `json:"-" bson:"_id,omitempty"`
	Name        string    `json:"name"`
	Path        string    `json:"path"`
	CreatedTime time.Time `json:"created_time"`
}

func (p *Playbook) Base() *ansible.BasePlaybook {
	return &ansible.BasePlaybook{
		Path: p.Path,
	}
}
