package models

import "time"

const (
	TaskStateFinished   = "finished"
	TaskStateRunning    = "running"
	TaskStatePending    = "pending"
	TaskStateScheduling = "scheduling"
)

type Task struct {
	Id           string     `json:"-" bson:"_id,omitempty"`
	Uid          string     `json:"uid" bson:"uid"`
	State        string     `json:"state" bson:"state"`
	Success      bool       `json:"success" bson:"success"`
	WebSocket    string     `json:"web_socket" bson:"wob_socket"`
	Scheduled    bool       `json:"scheduled" bson:"scheduled"`
	Worker       string     `json:"worker" bson:"worker"`
	CreatedTime  *time.Time `json:"created_time" bson:"created_time"`
	ScheduleTime *time.Time `json:"schedule_time" bson:"schedule_time"`
	Args         string     `json:"args",bson:"args"`
	Result       Result     `json:"result" bson:"result"`
}
type Result struct {
	StartTime *time.Time `json:"start_time" bson:"start_time"`
	EndTime   *time.Time `json:"end_time" bson:"end_time"`
	ExitCode  int        `json:"exit_code" bson:"exit_code"`
	Message   string     `json:"message" bson:"message"`
	Finished  bool       `json:"finished" bson:"finished"`
}
