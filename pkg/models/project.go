package models

type Playbook string

type Project struct {
	Name      string
	Playbooks []Playbook
}
