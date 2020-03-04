package worker

import (
	"kobe/pkg/models"
	"os"
)

type Runnable interface {
	Run(args map[string]string, workPath string, logFile *os.File, result *models.Result)
}
