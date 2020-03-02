package models

type PlaybookSet struct {
	Name      string `json:"name"`
	Path      string `json:"path"`
	Playbooks []Playbook `json:"playbooks"`
}

type Playbook struct {
	Name string `json:"name"`
}
