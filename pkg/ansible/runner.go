package ansible

import (
	ansibler "github.com/apenella/go-ansible"
	"kobe/pkg/models"
	"os"
	"path"
	"time"
)

type PlaybookRunner struct {
	Inventory string
	Playbook  string
	Options   map[string]interface{}
}

func (p *PlaybookRunner) Run(workPath string, logFile *os.File, result *models.Result) {
	pwd, _ := os.Getwd()
	inventoryPath := path.Join(pwd, "data", "inventory", p.Inventory)
	playbookPath := path.Join(pwd, "data", "playbook", p.Playbook)
	_ = os.Chdir(workPath)
	defer os.Chdir(pwd)
	pb := &ansibler.AnsiblePlaybookCmd{
		Playbook:   playbookPath,
		ExecPrefix: "kobe",
		Options: &ansibler.AnsiblePlaybookOptions{
			Inventory: inventoryPath,
		},
		Writer: logFile,
	}
	if err := pb.Run(); err != nil {
		result.Success = false
		result.Message = err.Error()
	} else {
		result.Success = true
	}
	result.EndTime = time.Now()
}
