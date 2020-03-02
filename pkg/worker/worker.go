package worker

import (
	"encoding/json"
	"fmt"
	"kobe/pkg/connections"
	"kobe/pkg/logger"
	"kobe/pkg/models"
	"os"
	"time"
)

var log = logger.Logger

type Worker struct {
	CurrentTask *models.Task
}

const (
	taskQueueKey = "queue"
)

func (w *Worker) Listen() {
	log.Infof("worker: %s started ")
	for {
		log.Info("waiting for task...")
		var task models.Task
		taskUid := connections.Redis.BRPop(-1, taskQueueKey).String()
		log.Infof("received a task :%s", task.Uid)
		taskJson := connections.Redis.HGetAll(taskUid).String()
		if err := json.Unmarshal([]byte(taskJson), &task); err != nil {
			log.Errorf("invalid message, can not parse json to task reason", err)
			continue
		}
		w.CurrentTask = &task
	}
}

type Runnable interface {
	Run(args map[string]interface{}, workPath string, stdout *os.File) (models.Result, error)
}

func (w *Worker) work()  {
	w.CurrentTask.State = models.TaskStateRunning
	w.saveTask()
	// runner 所需要参数  playbook inventory args workPath logfile
	time.Sleep(10 * time.Second)
	fmt.Println("handle task success")
	w.CurrentTask.State = models.TaskStateFinished
}

func (w *Worker) saveTask() {
	connections.Redis.HSet(w.CurrentTask.Uid, w.CurrentTask)
}

func Run() {
	connections.ConnectRedis()
	defer connections.Redis.Close()
	c := Worker{
		CurrentTask: nil,
	}
	c.Listen()
}
