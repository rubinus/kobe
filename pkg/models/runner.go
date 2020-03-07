package models

type CommonRunRequest struct {
	Inventory string `json:"inventory"`
}

type ImRunRequest struct {
	Inventory Inventory `json:"inventory"`
}

type ImRunPlaybookRequest struct {
	Dir      string `json:"dir"`
	Playbook string `json:"playbook"`
	ImRunRequest
}

type ImRunAdhocRequest struct {
	Pattern string `json:"pattern"`
	Module  string `json:"module"`
	Arg     string `json:"arg"`
	ImRunRequest
}

type RunPlaybookRequest struct {
	Dir      string `json:"dir"`
	Playbook string `json:"playbook"`
	CommonRunRequest
}

type RunAdhocRequest struct {
	Pattern string `json:"pattern"`
	Module  string `json:"module"`
	Arg     string `json:"arg"`
	CommonRunRequest
}
