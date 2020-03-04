package models

type CommonRunRequest struct {
	Inventory string `json:"inventory"`
}

type RunPlaybookRequest struct {
	Playbook string `json:"playbook"`
	CommonRunRequest
}

type RunAdhocRequest struct {
	Pattern   string `json:"pattern"`
	Module string `json:"module"`
	Arg    string `json:"arg"`
	CommonRunRequest
}

