package logger

import (
	"github.com/apex/log"
	"time"
)

var Logger *log.Entry

func init() {
	log.SetHandler(Default)
	log.SetLevel(log.InfoLevel)
	Logger = log.WithFields(log.Fields{
		"timestamp": time.Now().Format("2006-01-02 15:04:05"),
	})
}
