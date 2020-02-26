package task

import (
	"kobe/pkg/ansible"
	"kobe/pkg/model/inventory"
	"kobe/pkg/model/playbook"
	"os"
	"path"
	"time"
)

const basePath = "data/task/"

type Result struct {
	StartTime time.Time
	EndTime   time.Time
	ExitCode  int
	Run       bool
	Success   bool
	Message   string
}

type Task struct {
	Id        string
	Inventory *inventory.Inventory
}

type PlaybookTask struct {
	*Task
	Playbook *playbook.Playbook
}

func (t *PlaybookTask) GetRunner() (*ansible.PlaybookRunner, error) {
	wp, err := t.InitWorkPath()
	if err != nil {
		return nil, err
	}
	logFile, err := t.InitLogFile(wp)
	if err != nil {
		return nil, err
	}
	return ansible.NewPlaybookRunner(t.Playbook.BasePlaybook,
		t.Inventory.BaseInventory, wp, logFile), nil
}

func (t Task) InitWorkPath() (string, error) {
	p := path.Join(basePath, t.Id)
	if err := os.Mkdir(p, 0755); err != nil {
		return "", err
	}
	return p, nil
}

func (t Task) InitLogFile(workPath string) (*os.File, error) {
	p := path.Join(workPath, "task.log")
	if _, err := os.Stat(p); err != nil {
		file, err := os.Create(p)
		if err != nil {
			return nil, err
		}
		return file, nil
	}
	file, err := os.Open(p)
	if err != nil {
		return nil, err
	}
	return file, nil
}
