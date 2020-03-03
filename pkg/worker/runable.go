package worker

import (
	"kobe/pkg/models"
	"os"
)

type Runnable interface {
	Run(args map[string]interface{}, workPath string, logFile *os.File, result *models.Result)
}
