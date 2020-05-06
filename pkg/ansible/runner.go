package ansible

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"kobe/pkg/models"
	"os"
	"os/exec"
	"path"
	"text/template"
	"time"
)

const (
	ansibleTemplateFileName = "ansible.cfg.tmpl"
	ansibleCfgFileName      = "ansible.cfg"
	ansiblePluginDirName    = "plugins"
	resultFileName          = "result.json"
	tempPath                = "tmp"
)

type PlaybookRunner struct {
	Playbook models.Playbook
}

func (p *PlaybookRunner) Run(inventoryId, play string, result *models.Result) {

	workPath, err := initWorkSpace(p.Playbook)
	if err != nil {
		result.Message = err.Error()
		return
	}
	pwd, err := os.Getwd()
	if err != nil {
		result.Message = err.Error()
		return
	}

	os.Chdir(workPath)
	defer func() {
		os.Chdir(pwd)
		result.EndTime = time.Now()
	}()
	cmd := exec.Command("ansible",
		"-i", fmt.Sprintf(path.Join(pwd, "kobe")), fmt.Sprintf("--%s", inventoryId))
	if err := cmd.Run(); err != nil {
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
		return nil, err
	}
	var result interface{}
	if err := json.Unmarshal(content, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func initWorkSpace(playbook models.Playbook) (string, error) {
	workPath := path.Join(tempPath, string(playbook))
	if err := os.MkdirAll(workPath, 0755); err != nil {
		return "", err
	}
	if err := renderConfig(workPath); err != nil {
		return "", err
	}

	if err := initPlugin(workPath); err != nil {
		return "", err
	}
	return workPath, nil
}

func renderConfig(workPath string) error {
	tmpl := path.Join(workPath, "ansible", ansibleTemplateFileName)
	file, err := os.OpenFile(path.Join(workPath, ansibleCfgFileName), os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		return err
	}
	defer file.Close()
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		return err
	}
	data := map[string]interface{}{}
	if err := t.Execute(file, data); err != nil {
		return err
	}
	return nil
}

func initPlugin(workPath string) error {
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}
	if err := os.Symlink(path.Join(pwd, "ansible", ansiblePluginDirName),
		path.Join(workPath, ansiblePluginDirName)); err != nil {
		return err
	}
	return nil
}
