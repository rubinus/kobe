package models

import "time"

const (
	TaskStateFinished = "finished"
	TaskStateRunning  = "running"
	TaskStatePending  = "pending"
)

type Task struct {
	Uid         string            `json:"uid" bson:"uid"`
	State       string            `json:"state" bson:"state"`
	CreatedTime time.Time         `json:"created_time" bson:"created_time"`
	Args        map[string]string `json:"args",bson:"args"`
	*Result
}
type Result struct {
	StartTime time.Time `json:"start_time" bson:"start_time"`
	EndTime   time.Time `json:"end_time" bson:"end_time"`
	ExitCode  int       `json:"exit_code" bson:"exit_code"`
	Message   string    `json:"message" bson:"message"`
	Finished  bool      `json:"finished" bson:"finished"`
	Success   bool      `json:"success" bson:"success"`
	WebSocket string    `json:"web_socket" bson:"wob_socket"`
}
