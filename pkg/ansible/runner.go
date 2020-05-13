package ansible

import (
	"fmt"
	"github.com/prometheus/common/log"
	"io"
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
)

type PlaybookRunner struct {
	Project  api.Project
	Playbook string
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

	if _, err := exec.LookPath(constant.AnsiblePlaybookBinPath); err != nil {
		result.Success = false
		result.Message = err.Error()
		log.Error(err)
		return
	}
	cmd := exec.Command(constant.AnsiblePlaybookBinPath,
		"-i", constant.InventoryProviderBinPath,
		path.Join(constant.ProjectDir, p.Project.Name, p.Playbook))
	cmdEnv := make([]string, 0)
	cmdEnv = append(cmdEnv, fmt.Sprintf("%s=%s", constant.TaskEnvKey, result.Id))
	cmd.Env = append(os.Environ(), cmdEnv...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		result.Success = false
		result.Message = err.Error()
		return
	}
	if err := cmd.Start(); err != nil {
		result.Success = false
		result.Message = err.Error()
		return
	}
	buf := make([]byte, 4096)
	for {
		nr, err := stdout.Read(buf)
		if nr > 0 {
			select {
			case ch <- buf[:nr]:
			default:
			}
		}
		if err != nil || io.EOF == err {
			break
		}
	}
	close(ch)
	if err = cmd.Wait(); err != nil {
		result.Success = false
		result.Message = err.Error()
		return
	}
	result.Success = true
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
