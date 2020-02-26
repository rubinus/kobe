package worker

import (
	"github.com/go-redis/redis/v7"
	"github.com/spf13/cobra/cobra/cmd"
	b "kobe/pkg/broker"
	"kobe/pkg/worker/task"
	"log"
)
import _ "kobe/pkg/broker"

type Worker struct {
	Id     string
	broker b.Broker
}

func (w *Worker) Run() {
	w.broker = b.Broker{
		Client: redis.NewClient(&redis.Options{
			Addr: "localhost:6379",
			DB:   0,
		}),
	}
	for {
		t, err := w.broker.GetTask()
		if err != nil {
			log.Println(err)
		}
		w.Handler(t)
	}

}

func PlaybookHandler(t *task.PlaybookTask) *task.Result {
	r, err := t.GetRunner()
	if err != nil {
		log.Println(err)
	}
	state := make(chan int)
	pid, err := r.Run()
	go getState(pid)
	<-state
}

func getState(pid string, chan int)  {
}
