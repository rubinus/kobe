package server

import (
	"errors"
	"fmt"
	"kobe/api"
	"kobe/pkg/ansible"
)

var InventoryCache = NewInventoryCache()

type RunnerManager struct{}

func (rm *RunnerManager) CreatePlaybookRunner(projectName, playbookName string, inventory *api.Inventory) (*ansible.PlaybookRunner, error) {
	err := preRunPlaybook(projectName, playbookName)
	if err != nil {
		return nil, err
	}
	id := InventoryCache.Put(inventory)
	pm := ProjectManager{}
	p, err := pm.GetProject(projectName)
	if err != nil {
		return nil, err
	}
	return &ansible.PlaybookRunner{
		Project:     *p,
		Playbook:    playbookName,
		InventoryId: id,
	}, nil
}

func preRunPlaybook(projectName, playbookName string) error {
	pm := ProjectManager{}
	p, err := pm.GetProject(projectName)
	if err != nil {
		return err
	}
	exists := false
	for _, playbook := range p.Playbooks {
		if playbook == playbookName {
			exists = true
		}
	}
	if !exists {
		return errors.New(fmt.Sprintf("can not find playbook:%s in project:%s", playbookName, projectName))
	}
	return nil
}
