package models

import (
	"kobe/pkg/logger"
	"time"
)

const (
	StateWorkerOnline  = "online"
	StateWorkerOffline = "offline"
)

var log = logger.Logger

type Worker struct {
	Id                  string    `json:"id" bson:"_id,omitempty"`
	Uid                 string    `json:"uid" bson:"uid"`
	State               string    `json:"state"`
	LastHealthCheckTime time.Time `json:"last_health_check_time"`
}
