package models

import (
	"encoding/json"
	"time"
)

const (
	TaskStateFinished = "finished"
	TaskStateRunning  = "running"
	TaskStatePending  = "pending"
)

type Task struct {
	Uid         string                 `json:"uid" bson:"uid"`
	State       string                 `json:"state" bson:"state"`
	CreatedTime time.Time              `json:"created_time"`
	Args        map[string]interface{} `json:"args"`
	Finished    bool                   `json:"finished"`
}

func (t Task) MarshalBinary() (data []byte, err error) {
	return json.Marshal(t)
}

type Result struct {
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time" `
	Message   string    `json:"message"`
	Success   bool      `json:"success"`
	Stdout    string    `json:"stdout"`
	Logfile   string    `json:"logfile"`
}

func (r Result) MarshalBinary() (data []byte, err error) {
	return json.Marshal(r)
}
