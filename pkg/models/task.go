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

type TaskRequest struct {
	Args map[string]string `json:"args"`
}

type Task struct {
	Uid         string            `json:"uid" bson:"uid"`
	Inventory   Inventory         `json:"inventory"`
	State       string            `json:"state" bson:"state"`
	CreatedTime time.Time         `json:"created_time"`
	Args        map[string]string `json:"args"`
	Finished    bool              `json:"finished"`
	Type        string            `json:"type"`
}

func (t Task) MarshalBinary() (data []byte, err error) {
	return json.Marshal(t)
}

type Result struct {
	StartTime time.Time   `json:"start_time"`
	EndTime   time.Time   `json:"end_time" `
	Message   string      `json:"message"`
	Success   bool        `json:"success"`
	Stdout    string      `json:"stdout"`
	Content   interface{} `json:"content"`
	Logfile   string      `json:"logfile"`
}

func (r Result) MarshalBinary() (data []byte, err error) {
	return json.Marshal(r)
}
