package ansible

import (
	ansibler "github.com/apenella/go-ansible"
	"github.com/apex/log"
	"kobe/pkg/models"
	"os"
	"path"
	"time"
)

type AdhocRunner struct{}

func (a *AdhocRunner) Run(args map[string]string, workPath string, logFile *os.File, result *models.Result) {
	pwd, _ := os.Getwd()
	if err := os.Chdir(workPath); err != nil {
		log.Errorf("can not chdir %s reason %s", workPath, err.Error())
	}
	defer os.Chdir(pwd)
	executer := AnsibleExecuter{
		Write: logFile,
	}
	inventoryPath := path.Join(pwd, "data", "inventory", args["inventory"])

	as := []string{
		"-i", inventoryPath, args["pattern"],
		"-m", args["module"],
		"-a", args["arg"],
	}
	if err := executer.Execute("ansible", as, ""); err != nil {
		result.Success = false
		result.Message = err.Error()
		log.Error(err.Error())
	} else {
		result.Success = true
	}

}

type PlaybookRunner struct{}

func (p *PlaybookRunner) Run(args map[string]string, workPath string, logFile *os.File, result *models.Result) {
	pwd, _ := os.Getwd()
	inventoryPath := path.Join(pwd, "data", "inventory", args["inventory"])
	playbookPath := path.Join(pwd, "data", "playbooks", args["dir"], args["playbook"])
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
		Exec: &AnsibleExecuter{
			Write: logFile,
		},
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
