package runner

import (
    "io"
    "kobe/pkg/ansible/inventory"
    "kobe/pkg/ansible/playbook"
    "os/exec"
)

const (
    ansiblePlaybookName = "ansible-playbook"
)

type PlaybookRunner struct {
    Inventory inventory.Inventory
    Playbook  playbook.Playbook
    Stdout    io.Writer
}

func (pr *PlaybookRunner) Run() error {
    cmd := exec.Command(ansiblePlaybookName)
    cmd.Stdout = pr.Stdout
    cmd.Stderr = pr.Stdout
    if err := cmd.Start(); err != nil {
        return err
    }
    if err := cmd.Wait(); err != nil {
        return err
    }
    return nil
}
