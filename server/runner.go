package server

import (
	"errors"
	"fmt"
	"github.com/patrickmn/go-cache"
	uuid "github.com/satori/go.uuid"
	"kobe/api"
	"kobe/pkg/ansible"
)

type RunnerManager struct {
	inventoryCache *cache.Cache
}

func (rm *RunnerManager) CreatePlaybookRunner(projectName, playbookName string, inventory *api.Inventory) (*ansible.PlaybookRunner, error) {
	err := preRunPlaybook(projectName, playbookName)
	if err != nil {
		return nil, err
	}
	id := uuid.NewV4().String()
	fmt.Println("inventory" + id)
	rm.inventoryCache.Set(id, inventory, cache.DefaultExpiration)
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
