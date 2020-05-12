package constant

import (
	"path"
)

const (
	InventoryProviderBinPath = "/Users/shenchenyang/go/bin/inventory"
	AnsiblePlaybookBinPath   = "ansible-playbook"
	TaskEnvKey               = "KO_TASK_ID"
)

var (
	BaseDir                 = "/Users/shenchenyang/go/src/kobe/"
	DataDir                 = path.Join(BaseDir, "data")
	WorkDir                 = path.Join(BaseDir, "work")
	ProjectDir              = path.Join(DataDir, "project")
	AnsibleResDir           = path.Join(BaseDir, "ansible")
	AnsiblePluginDir        = path.Join(AnsibleResDir, "plugins")
	AnsibleTemplateFilePath = path.Join(AnsibleResDir, "ansible.cfg.tmpl")
)
