package ansible

import (
	"github.com/chenhg5/collection"
	"io"
	"kobe/pkg/util"
	"os"
	"os/exec"
	"path"
	"sync"
)

var connectionOptions = []string{
	"ask-pass", "private-key", "user", "connection", "timeout", "ssh-common-args",
	"sftp-extra-args", "scp-extra-args", "ssh-extra-args",
}

var privilegeEscalationOptions = []string{
	"become", "become-method", "become-user", "ask-sudo-pass", "ask-su-pass", "ask-become-pass",
}

var playbookOptions = []string{
	"ask-vault-pass", "check", "diff", "extra-vars", "flush-cache", "force-handlers", "forks",
	"help", "inventorys", "inventorys-file", "limit", "list-hosts", "list-tags", "list-tasks", "module-path",
	"skip-tags", "start-at-task", "step", "syntax-check", "tags", "vault-id", "vault-password-file",
	"verbose", "version",
}

type runner struct {
	*BaseInventory
	WorkPath string
	Stdout   io.Writer
	Options  map[string]interface{}
	mutex    sync.Mutex
}

func (r runner) SetOption(name string, value interface{}) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.Options[name] = value
}

type PlaybookRunner struct {
	Playbook *BasePlaybook
	*runner
}

func NewPlaybookRunner(playbook *BasePlaybook, inventory *BaseInventory, workPath string, stdout io.Writer) *PlaybookRunner {
	return &PlaybookRunner{
		Playbook: playbook,
		runner: &runner{
			BaseInventory: inventory,
			WorkPath:      workPath,
			Options:       map[string]interface{}{},
			Stdout:        stdout,
		}}
}

func (p *PlaybookRunner) Run() (int, error) {
	currentPath, _ := os.Getwd()
	_ = os.Chdir(p.WorkPath)
	inventoryPath, _ := p.CreateInventory()
	allOptions := append(append(connectionOptions, privilegeEscalationOptions...), playbookOptions...)
	notSupportOptions := make([]string, 0)
	for option, _ := range p.Options {
		if !collection.Collect(allOptions).Has(option) {
			notSupportOptions = append(notSupportOptions, option)
			delete(p.Options, option)
			continue
		}
	}

	args := util.ArgsToStringArray(p.Options)
	args = append(args, "-i", inventoryPath)
	args = append(args, p.Playbook.Path)
	cmd := exec.Command("ansible-playbook", args...)
	cmd.Stdout = p.Stdout
	cmd.Stderr = p.Stdout
	if err := cmd.Start(); err != nil {
		return 0, err
	}
	_ = os.Chdir(currentPath)
	return cmd.ProcessState.Pid(), nil
}

func (p *PlaybookRunner) CreateInventory() (string, error) {
	inventoryPath := path.Join(p.WorkPath, "hosts.json")
	var f *os.File
	if _, err := os.Stat(inventoryPath); err != nil {
		f, _ = os.Create(inventoryPath)
	} else {
		f, _ = os.Open(inventoryPath)
	}
	defer f.Close()
	_, _ = f.WriteString(p.BaseInventory.String())
	return inventoryPath, nil
}
