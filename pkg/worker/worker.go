package worker

import (
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
)

func (w *Worker) Listen() {
	log.Infof("worker: %s started ")
	for {
		log.Info("waiting for task...")
		//var task models.Task
		taskUid := connections.Redis.BRPop(0, taskQueueKey).Val()[1]
		log.Infof("received a task :%s", taskUid)
		taskJson := connections.Redis.HGetAll(taskUid)
		log.Debugf("taskInfo:", taskJson.String())
		//if err := json.Unmarshal([]byte(taskJson), &task); err != nil {
		//	log.Errorf("invalid message, can not parse json to task reason:", err.Error())
		//	continue
		//}
		//w.CurrentTask = &task
	}
}

func (w *Worker) work() {
	w.CurrentTask.State = models.TaskStateRunning
	w.saveTask()
	result := models.Result{
		StartTime: time.Now(),
	}
	pwd, _ := os.Getwd()
	workPath := fmt.Sprintf(path.Join(pwd, w.CurrentTask.Uid))
	logFile, _ := os.Create(fmt.Sprintf("%s.log", path.Join(workPath, "log", w.CurrentTask.Uid)))
	_ = os.MkdirAll(workPath, 0755)
	runner := ansible.PlaybookRunner{
		Inventory: w.CurrentTask.Args["inventory"].(string),
		Playbook:  w.CurrentTask.Args["playbook"].(string),
	}
	runner.Run(workPath, logFile, &result)

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
