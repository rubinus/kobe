package inventory

type Inventory struct {
	Groups []Group `json:"groups"`
	Hosts  []Host  `json:"hosts"`
}

func NewInventory(groups []Group, hosts []Host) *Inventory {
	return &Inventory{Groups: groups, Hosts: hosts}
}




