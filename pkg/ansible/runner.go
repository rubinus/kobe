package ansible

import (
	"fmt"
	ansibler "github.com/apenella/go-ansible"
	"github.com/apex/log"
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
	playbookPath := path.Join(pwd, "data", "playbooks", fmt.Sprintf("%s.yml", p.Playbook))
	if err := os.Chdir(workPath); err != nil {
		log.Errorf("can not chdir %s reason %s", workPath, err.Error())
	}
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
		log.Error(err.Error())
	} else {
		result.Success = true
	}
	result.EndTime = time.Now()
}
