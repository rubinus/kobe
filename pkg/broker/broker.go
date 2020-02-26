package broker

import (
	"encoding/json"
	"github.com/go-redis/redis/v7"
	"kobe/pkg/worker/task"
)

const taskQueue = "kobe_task"
const resultQueue = "kobe_task_result"

type Broker struct {
	Client *redis.Client
}

func (b *Broker) Ping() error {
	return b.Client.Ping().Err()
}

func (b *Broker) PushTask(t task.Task) error {
	msg, err := json.Marshal(t)
	if err != nil {
		return err
	}
	return b.Client.LPush(taskQueue, msg).Err()
}

func (b *Broker) GetTask() (*task.Task, error) {
	msg, err := b.Client.BRPop(0, taskQueue).Result()
	if err != nil {
		return nil, err
	}
	var t task.Task
	var bytes = []byte(msg[0])
	if err := json.Unmarshal(bytes, t); err != nil {
		return nil, err
	}
	return &t, nil
}

func (b *Broker) PushResult(msg string) error {
	return b.Client.LPush(resultQueue, msg).Err()
}

func (b *Broker) GetResult(msg string) ([]string, error) {
	return b.Client.BRPop(0, resultQueue).Result()
}
