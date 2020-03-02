package ansible

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"kobe/pkg/db"
	"kobe/pkg/logger"
	"kobe/pkg/models"
	"os"
	"os/exec"
	"path"
	"time"
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

var log = logger.Logger

type PlaybookRunner struct{}

func (pr *PlaybookRunner) Run(args map[string]interface{}, workPath string, stdout *os.File) (*models.Result, error) {
	pwd, err := os.Getwd()
	defer os.Chdir(pwd)
	if err != nil {
		log.Errorf("can not get pwd, reason: %s", err)
		return nil, err
	}
	if err := os.Chdir(workPath); err != nil {
		log.Errorf("can not chdir %s", err)
		return nil, err
	}
	inventory, playbook, options, err := handleArgs(args)
	if err != nil {
		log.Errorf("can not parse args", err.Error())
		return nil, err
	}
	line := []string{}
	line := append(line, "-i", inventory)
	for k, v := range options {
		line = append(line, fmt.Sprintf("--%s", k))
		line = append(line, v)
	}
	line = append(line, playbook.Path)

	cmd := exec.Command("ansible-playbook", line...)
	cmd.Stdout = stdout
	cmd.Stderr = stdout

	start := time.Now()
	result := models.Result{
		StartTime: &start,
		ExitCode:  0,
		Message:   "",
	}
	if err := cmd.Run(); err != nil {
		result.Message = err.Error()
		result.Finished = false
		return &result, nil
	}
	end := time.Now()
	result.Finished = true
	result.EndTime = &end
	result.ExitCode = cmd.ProcessState.ExitCode()
	return &result, nil
}
func handleArgs(args map[string]interface{}) (*BaseInventory, *BasePlaybook, map[string]string, error) {
	session := db.Session.Clone()
	d := session.DB(db.Mongo.Database)
	inventoryName := args["inventory"]
	playbookName := args["playbook"]
	var inventory models.Inventory
	var playbook models.Playbook
	if err := d.C("inventory").Find(bson.M{"name": inventoryName}).One(&inventory); err != nil {
		log.Errorf("can not find inventory %s reason %s", inventoryName, err)
		return nil, nil, nil, err
	}
	if err := d.C("playbook").Find(bson.M{"name": playbookName}).One(&playbook); err != nil {
		log.Errorf("can not find playbook %s reason %s", inventoryName, err)
		return nil, nil, nil, err
	}
	options := map[string]string{}
	_, ok := args["options"]
	if ok {
		options = args["options"].(map[string]string)
	} else {
		options = make(map[string]string)
	}
	baseInventory := inventory.Base()
	basePlaybook := playbook.Base()
	return baseInventory, basePlaybook, options, nil

}
