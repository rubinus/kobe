package ansible

import (
	"fmt"
	"io/ioutil"
	"kobe/api"
	"kobe/pkg/constant"
	"os"
	"os/exec"
	"path"
	"text/template"
	"time"
)

const (
	ansibleCfgFileName   = "ansible.cfg"
	ansiblePluginDirName = "plugins"
	resultFileName       = "result.json"
)

type PlaybookRunner struct {
	Project     api.Project
	Playbook    string
	InventoryId string
}

func (p *PlaybookRunner) Run(ch chan []byte, result *api.Result) {
	workPath, err := initWorkSpace(p.Project)
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
		result.EndTime = time.Now().String()
	}()
	cmd := exec.Command(
		constant.AnsiblePlaybookBinPath,
		"-i", constant.InventoryProviderBinPath,
		path.Join(constant.ProjectDir, p.Project.Name, p.Playbook))
	cmd.Env = append(os.Environ(), fmt.Sprintf("%s=%s", constant.InventoryEnvKey, p.InventoryId))
	reader, err := cmd.StdoutPipe()
	if err != nil {
		result.Success = false
		result.Message = err.Error()
		return
	}
	go func() {
		var buffer []byte
		for {
			_, err = reader.Read(buffer)
			if err != nil {
				break
			}
			ch <- buffer
		}
		close(ch)
	}()
	if err := cmd.Start(); err != nil {
		result.Success = false
		result.Message = err.Error()
		return
	}
	if err = cmd.Wait(); err != nil {
		result.Success = false
		result.Message = err.Error()
		return
	}
	_ = reader.Close()
	content, err := readResultFile(workPath)
	if err != nil {
		result.Success = false
		result.Message = err.Error()
		return
	}
	result.Content = content
	result.Success = true
}

func readResultFile(workPath string) (string, error) {
	p := path.Join(workPath, resultFileName)
	content, err := ioutil.ReadFile(p)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func initWorkSpace(project api.Project) (string, error) {
	workPath := path.Join(constant.WorkDir, project.Name)
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
	tmpl := constant.AnsibleTemplateFilePath
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
	projectPluginDir := path.Join(workPath, ansiblePluginDirName)
	_, err := os.Stat(projectPluginDir)
	if os.IsNotExist(err) {
		if err := os.Symlink(constant.AnsiblePluginDir, path.Join(workPath, ansiblePluginDirName))
			err != nil {
			return err
		}
		return nil
	}

	return err
}
