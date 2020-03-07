package ansible

import (
	"encoding/json"
	"errors"
	"fmt"
	ansibler "github.com/apenella/go-ansible"
	"io/ioutil"
	"kobe/pkg/logger"
	"kobe/pkg/models"
	"os"
	"path"
	"text/template"
)

const (
	inventoryFileName       = "hosts.json"
	ansibleTemplateFileName = "ansible.cfg.tmpl"
	ansibleCfgFileName      = "ansible.cfg"
	ansiblePluginDirName    = "plugins"
	resultFileName          = "result.json"
)

var log = logger.Logger

type AdhocRunner struct{}

func (a *AdhocRunner) Run(args map[string]string, inventory models.Inventory, workPath string, logFile *os.File, result *models.Result) {
	if err := initWorkSpace(workPath, inventory); err != nil {
		result.Success = false
		result.Message = err.Error()
		return
	}
	inventoryPath := path.Join(workPath, inventoryFileName)
	executor := AnsibleExecuter{
		Write: logFile,
	}
	as := []string{
		"-i", inventoryPath, args["pattern"],
		"-m", args["module"],
		"-a", args["arg"],
	}
	pwd, _ := os.Getwd()
	os.Chdir(workPath)
	defer os.Chdir(pwd)
	if err := executor.Execute("ansible", as, ""); err != nil {
		result.Success = false
		result.Message = err.Error()
	} else {
		content, err := readResultFile(workPath)
		if err != nil {
			result.Message = err.Error()
		}
		result.Content = content
		result.Success = true
	}
}

type PlaybookRunner struct{}

func (p *PlaybookRunner) Run(args map[string]string, inventory models.Inventory, workPath string, logFile *os.File, result *models.Result) {
	if err := initWorkSpace(workPath, inventory); err != nil {
		result.Success = false
		result.Message = err.Error()
		return
	}
	inventoryPath := path.Join(workPath, inventoryFileName)
	pb := &ansibler.AnsiblePlaybookCmd{
		Playbook: "",
		Options: &ansibler.AnsiblePlaybookOptions{
			Inventory: inventoryPath,
		},
		Exec: &AnsibleExecuter{
			Write: logFile,
		},
	}
	pwd, _ := os.Getwd()
	os.Chdir(workPath)
	defer os.Chdir(pwd)
	if err := pb.Run(); err != nil {
		result.Success = false
		result.Message = err.Error()
	} else {
		content, err := readResultFile(workPath)
		if err != nil {
			result.Message = err.Error()
		}
		result.Content = content
		result.Success = true
	}

}

func readResultFile(workPath string) (interface{}, error) {
	p := path.Join(workPath, resultFileName)
	content, err := ioutil.ReadFile(p)
	if err != nil {
		log.Errorf(fmt.Sprintf("read result error: %s", err.Error()))
		return "", errors.New(fmt.Sprintf("can not read result file %s", err.Error()))
	}
	var result interface{}
	if err := json.Unmarshal(content, &result); err != nil {
		log.Errorf(fmt.Sprintf("parse obj error: %s", err.Error()))
		return "", errors.New(fmt.Sprintf("parse obj error: %s", err.Error()))
	}

	return result, nil
}

func initWorkSpace(workPath string, inventory models.Inventory) error {
	if err := randerAnsibleConfig(workPath); err != nil {
		log.Errorf(fmt.Sprintf("rander config error: %s", err.Error()))
		return err
	}
	if err := writeInventoryFile(workPath, inventory); err != nil {
		log.Errorf(fmt.Sprintf("write inventory config error: %s", err.Error()))
		return err
	}
	if err := initAnsiblePlugin(workPath); err != nil {
		log.Errorf(fmt.Sprintf("init plugin error: %s", err.Error()))
		return err
	}
	return nil
}

func randerAnsibleConfig(workPath string) error {
	pwd, err := os.Getwd()
	if err != nil {
		return errors.New(fmt.Sprintf("can not get pwd reason %s", err))
	}
	tmpl := path.Join(pwd, "ansible", ansibleTemplateFileName)
	file, err := os.OpenFile(path.Join(workPath, ansibleCfgFileName), os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		return errors.New(fmt.Sprintf("can not create file reason %s", err.Error()))
	}
	defer file.Close()
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		return errors.New(fmt.Sprintf("can not parsefile %s reason %s", tmpl, err.Error()))
	}
	data := map[string]interface{}{}
	if err := t.Execute(file, data); err != nil {
		return errors.New(fmt.Sprintf("can not execute parse reason:%s", err.Error()))
	}
	return nil
}

func initAnsiblePlugin(workPath string) error {
	pwd, err := os.Getwd()
	if err != nil {
		return errors.New(fmt.Sprintf("can not get pwd reason %s", err))
	}
	if err := os.Symlink(path.Join(pwd, "ansible", ansiblePluginDirName), path.Join(workPath, ansiblePluginDirName)); err != nil {
		return errors.New(fmt.Sprintf("can not create symlink  reason %s", err))
	}
	return nil
}

func writeInventoryFile(workPath string, inventory models.Inventory) error {
	file, err := os.OpenFile(path.Join(workPath, inventoryFileName), os.O_CREATE, 0755)
	if err != nil {
		return errors.New(fmt.Sprintf("can not create file reason %s", err.Error()))
	}
	defer file.Close()
	data := inventory.Data()
	bs, err := json.Marshal(data)
	if err != nil {
		return errors.New(fmt.Sprintf("can not parse inventory to string reason %s", err.Error()))
	}
	if err := ioutil.WriteFile(file.Name(), bs, 0755); err != nil {
		return errors.New(fmt.Sprintf("can not write inventory file reason %s", err.Error()))
	}
	return nil
}
