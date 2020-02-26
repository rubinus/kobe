package inventory

import (
	"kobe/pkg/ansible"
	"kobe/pkg/model"
)

type Inventory struct {
	*model.Model
	*ansible.BaseInventory
}
