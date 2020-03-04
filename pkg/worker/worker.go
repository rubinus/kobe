package worker

import (
	"encoding/json"
	"fmt"
	"kobe/pkg/ansible"
	"kobe/pkg/connections"
	"kobe/pkg/logger"
	"kobe/pkg/models"
	"os"
	"path"
	"time"
)

var log = logger.Logger

type Worker struct {
	CurrentTask *models.Task
}

const (
	taskQueueKey = "queue"
	taskSetKey   = "task"
	resultKey    = "result"
)

func (w *Worker) Listen() {
	log.Infof("worker: %s started ")
	for {
		log.Info("waiting for task...")
		var task models.Task
		taskUid := connections.Redis.BRPop(0, taskQueueKey).Val()[1]
		log.Infof("received a task :%s", taskUid)
		taskJson, err := connections.Redis.HGet(taskSetKey, taskUid).Result()
		if err != nil {
			log.Errorf("can not read task info :", err.Error())
			continue
		}
		if err := json.Unmarshal([]byte(taskJson), &task); err != nil {
			log.Errorf("invalid message, can not parse json to task reason:", err.Error())
			continue
		}
		w.CurrentTask = &task
		w.work()
	}
}

func (w *Worker) work() error {
	w.CurrentTask.State = models.TaskStateRunning
	pwd, _ := os.Getwd()
	workPath := fmt.Sprintf(path.Join(pwd, "data", "task", w.CurrentTask.Uid))
	if err := os.MkdirAll(workPath, 0755); err != nil {
		log.Errorf("can not work dir %s reason %s", workPath, err.Error())
		return err
	}
	logPath := fmt.Sprintf("%s.log", path.Join(workPath, w.CurrentTask.Uid))
	logFile, err := os.Create(logPath)
	if err := w.saveTask(); err != nil {
		log.Errorf("can not save task reason: %s", err.Error())
		return err
	}
	result := models.Result{
		StartTime: time.Now(),
		Logfile:   logPath,
	}
	if err != nil {
		log.Errorf("can not create log file reason %s", err.Error())
	}
	runner := ansible.PlaybookRunner{
		Inventory: w.CurrentTask.Args["inventory"].(string),
		Playbook:  w.CurrentTask.Args["playbook"].(string),
	}
	runner.Run(workPath, logFile, &result)
	w.CurrentTask.State = models.TaskStateFinished
	if err := w.saveTask(); err != nil {
		log.Errorf("can not save task reason: %s", err.Error())
		return err
	}
	if err := w.saveResult(result); err != nil {
		log.Errorf("can not save task result reason: %s", err.Error())
	}
	return nil
}

func (w *Worker) saveTask() error {
	_, err := connections.Redis.HSet(taskSetKey, w.CurrentTask.Uid, w.CurrentTask).Result()
	return err
}

func (w *Worker) saveResult(result models.Result) error {
	_, err := connections.Redis.HSet(resultKey, w.CurrentTask.Uid, result).Result()
	return err
}

func Run() {
	connections.ConnectRedis()
	defer connections.Redis.Close()
	c := Worker{
		CurrentTask: nil,
	}
	c.Listen()
}
