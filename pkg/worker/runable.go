package worker

import (
	"kobe/pkg/models"
	"os"
)

type Runnable interface {
	Run(args map[string]string, inventory models.Inventory, workPath string, logFile *os.File, result *models.Result)
}
