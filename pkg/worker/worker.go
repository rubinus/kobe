package container

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"kobe/pkg/db"
	"kobe/pkg/logger"
	"kobe/pkg/models"
	"os"
	"path"
	"time"
)

var log = logger.Logger

type Container struct {
	worker      *models.Worker
	Db          *mgo.Database
	CurrentTask *models.Task
}

func (w *Container) Listen() {
	uid, err := readOrCreateWorkerUId()
	if err != nil {
		log.Fatalf("can not get worker Uid:  %s", err)
	}
	w.worker.Uid = uid
	if err := w.save(); err != nil {
		log.Fatalf("can not start worker: %s", err)
	}
	log.Infof("worker: %s started ", w.worker.Uid)
	go w.SendHealth()
	for {
		log.Info("waiting for task...")
		time.Sleep(5 * time.Second)
		c := w.Db.C("task")
		var task models.Task
		if err := c.Find(bson.M{"state": models.TaskStateScheduling}).Sort("-created_time").One(&task); err != nil {
			if err == mgo.ErrNotFound {
				log.Info("no task in queue")
			} else {
				log.Error(err.Error())
			}
			continue
		}
		if err := w.schedule(&task); err != nil {
			log.Error(err.Error())
			continue
		}
		if err := w.work(); err != nil {
			log.Error(err.Error())
		}
	}
}

func (w *Container) work() error {
	w.CurrentTask.State = models.TaskStateRunning
	w.CurrentTask.StartTime = time.Now()
	_ = w.saveTask()
	p := w.CurrentTask.Playbook
	i := w.CurrentTask.Inventory
	fmt.Printf("handle task %s %s", p, i)
	time.Sleep(10 * time.Second)
	fmt.Println("handle task success")
	w.CurrentTask.State = models.TaskStateFinished
	w.CurrentTask.EndTime = time.Now()
	return w.saveTask()
}

func (w *Container) schedule(task *models.Task) error {
	w.CurrentTask = task
	w.CurrentTask.Scheduled = true
	w.CurrentTask.ScheduleTime = time.Now()
	w.CurrentTask.Worker = w.worker.Uid
	w.CurrentTask.State = models.TaskStatePending
	return w.saveTask()
}

func (w *Container) saveTask() error {
	c := w.Db.C("task")
	w.CurrentTask.Id = ""
	if err := c.Update(bson.M{"uid": w.CurrentTask.Uid}, &w.CurrentTask); err != nil {
		return err
	}
	return nil
}

func readOrCreateWorkerUId() (string, error) {
	pwd, _ := os.Getwd()
	idPath := path.Join(pwd, "data", "worker", "uid")
	if _, err := os.Stat(idPath); err != nil {
		if os.IsPermission(err) {
			return "", err
		}
		if !os.IsExist(err) {
			uid := uuid.NewV4().String()
			if err := writeWorkerUid(uid); err != nil {
				return "", err
			}
			return uid, nil
		}
	}
	file, err := os.OpenFile(idPath, os.O_RDONLY, 0755)
	if err != nil {
		return "", err
	}
	bs, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	if string(bs) == "" {
		uid := uuid.NewV4().String()
		if err := writeWorkerUid(uid); err != nil {
			return "", err
		}
		return uid, nil
	}
	return string(bs), nil
}

func writeWorkerUid(uid string) error {
	pwd, _ := os.Getwd()
	workerPath := path.Join(pwd, "data", "worker")
	idPath := path.Join(workerPath, "uid")
	if err := os.MkdirAll(workerPath, 0754); err != nil {
		return err
	}
	if err := ioutil.WriteFile(idPath, []byte(uid), 0755); err != nil {
		return err
	}
	return nil
}

func (w *Container) SendHealth() {
	for {
		time.Sleep(5 * time.Second)
		w.worker.State = models.StateWorkerOnline
		w.worker.LastHealthCheckTime = time.Now()
		if err := w.save(); err != nil {
			log.Warnf("can not send health: %s", err)
		}
	}
}

func (w *Container) save() error {
	if _, err := w.Db.C("worker").Upsert(bson.M{"uid": w.worker.Uid}, w.worker); err != nil {
		return err
	}
	return nil
}

func Run() {
	db.Connect()
	s := db.Session.Clone()
	defer s.Close()
	d := s.DB(db.Mongo.Database)
	c := Container{
		worker: &models.Worker{
			State: models.StateWorkerOnline,
		},
		Db:          d,
		CurrentTask: nil,
	}
	c.Listen()
}
