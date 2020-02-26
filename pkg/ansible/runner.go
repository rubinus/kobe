package ansible

import (
    "github.com/chenhg5/collection"
    "kobe/pkg/util"
    "os"
    "os/exec"
    "path"
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

type Runnable interface {
    Run() error
}

type PlaybookRunner struct {
    *BasePlaybook
    *BaseInventory
    WorkPath string
    Options  map[string]interface{}
}

func (p *PlaybookRunner) Run(callback *CallBack) error {
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
    args = append(args, p.BasePlaybook.Path)
    cmd := exec.Command("ansible-playbook", args...)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    if err := cmd.Run(); err != nil {
        return err
    }
    if cmd.ProcessState.ExitCode() == 0 {
        callback.onSuccess()
    } else {
        callback.onError()
    }
    _ = os.Chdir(currentPath)
    return nil
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
