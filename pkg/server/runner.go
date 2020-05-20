package server

import (
	"errors"
	"fmt"
	"github.com/patrickmn/go-cache"
	"github.com/KubeOperator/kobe/pkg/ansible"
)

type RunnerManager struct {
	inventoryCache *cache.Cache
}

func (rm *RunnerManager) CreatePlaybookRunner(projectName, playbookName string) (*ansible.PlaybookRunner, error) {
	err := preRunPlaybook(projectName, playbookName)
	if err != nil {
		return nil, err
	}
	pm := ProjectManager{}
	p, err := pm.GetProject(projectName)
	if err != nil {
		return nil, err
	}
	return &ansible.PlaybookRunner{
		Project:  *p,
		Playbook: playbookName,
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
