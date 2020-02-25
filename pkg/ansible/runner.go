package ansible

import (
    "github.com/chenhg5/collection"
    "kobe/pkg/util"
    "os"
    "os/exec"
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
    "help", "inventory", "inventory-file", "limit", "list-hosts", "list-tags", "list-tasks", "module-path",
    "skip-tags", "start-at-task", "step", "syntax-check", "tags", "vault-id", "vault-password-file",
    "verbose", "version",
}

type Runnable interface {
    Run() error
}

type PlaybookRunner struct {
    Playbook  BasePlaybook
    Inventory BaseInventory
    Options   map[string]interface{}
}

func (p *PlaybookRunner) Run(callback *CallBack) error {
    allOptions := append(append(connectionOptions, privilegeEscalationOptions...), playbookOptions...)
    notSupportOptions := make([]string, 10)
    for option, _ := range p.Options {
        if !collection.Collect(allOptions).Has(option) {
            notSupportOptions = append(notSupportOptions, option)
            delete(p.Options, option)
            continue
        }
    }
    args := util.ArgsToStringArray(p.Options)
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
    return nil
}
