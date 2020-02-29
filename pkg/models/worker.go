package models

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"kobe/pkg/logger"
	"time"
)

const (
	StateWorkerOnline  = "online"
	StateWorkerOffline = "offline"
)

var log = logger.Logger

type Worker struct {
	Id                  string    `json:"id" bson:"_id"`
	State               string    `json:"state"`
	LastHealthCheckTime time.Time `json:"last_health_check_time"`
	Db                  *mgo.Database
	CurrentTask         Task
}

func (w *Worker) Listen() {
	log.Infof("worker: %s started ", w.Id)
	for {
		time.Sleep(10 * time.Second)
		w.SendHealth()
		log.Info("waiting for task...")
		c := w.Db.C("task")
		var task Task
		_ = c.Find(nil).Sort().One(&task)
		err := w.schedule(task)
		if err != nil {
			log.Warn(err.Error())
			continue
		}
	}
}

func (w *Worker) work() {
	w.CurrentTask.State = StateRunning
	w.CurrentTask.StartTime = time.Now()
	_ = w.saveTask()
	p := w.CurrentTask.Playbook
	i := w.CurrentTask.Inventory
	fmt.Printf("handle task %s %s", p, i)
	time.Sleep(10 * time.Second)
	fmt.Println("handle task success")
	w.CurrentTask.State = StateFinished
	w.CurrentTask.EndTime = time.Now()
	_ = w.saveTask()
}

func (w *Worker) schedule(task Task) error {
	w.CurrentTask = task
	w.CurrentTask.Scheduled = true
	w.CurrentTask.ScheduleTime = time.Now()
	w.CurrentTask.Worker = w.Id
	return w.saveTask()
}

func (w *Worker) saveTask() error {
	c := w.Db.C("task")
	if err := c.Update(bson.M{"id": w.CurrentTask.Id}, &w.CurrentTask); err != nil {
		return err
	}
	return nil
}

func (w *Worker) SendHealth() {
	go func() {
		w.State = StateWorkerOnline
		w.LastHealthCheckTime = time.Now()
		_, _ = w.Db.C("worker").Upsert(bson.M{"_id": w.Id}, w)
	}()
}
