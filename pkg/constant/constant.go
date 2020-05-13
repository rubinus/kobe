package constant

import (
	"github.com/spf13/viper"
	"path"
)

const (
	InventoryProviderBinPath = "kobe-inventory"
	AnsiblePlaybookBinPath   = "ansible-playbook"
	TaskEnvKey               = "KO_TASK_ID"
)

var (
	BaseDir                 = "/var/kobe"
	LibDir                  = path.Join(BaseDir, "lib")
	DataDir                 = path.Join(BaseDir, "data")
	WorkDir                 = path.Join(BaseDir, "work")
	AnsibleLibDir           = path.Join(LibDir, "ansible")
	ProjectDir              = path.Join(DataDir, "project")
	AnsiblePluginDir        = path.Join(AnsibleLibDir, "plugins")
	AnsibleTemplateFilePath = path.Join(AnsibleLibDir, "ansible.cfg.tmpl")
)

func Init() {
	BaseDir = viper.GetString("base")
}
