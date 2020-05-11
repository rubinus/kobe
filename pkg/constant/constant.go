package constant

import "path"

const (
	BaseDir                  = "/Users/shenchenyang/go/src/kobe/"
	InventoryProviderBinPath = "/Users/shenchenyang/go/bin/inventory"
	AnsiblePlaybookBinPath   = "ansible-playbook"
	InventoryEnvKey          = "KO_INVENTORY_ID"
)

var (
	DataDir                 = path.Join(BaseDir, "data")
	WorkDir                 = path.Join(BaseDir, "work")
	ProjectDir              = path.Join(DataDir, "project")
	AnsibleResDir           = path.Join(BaseDir, "ansible")
	AnsiblePluginDir        = path.Join(AnsibleResDir, "plugins")
	AnsibleTemplateFilePath = path.Join(AnsibleResDir, "ansible.cfg.tmpl")
)
