package broker

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v7"
	"kobe/pkg/worker/task"
	"log"
	"testing"
)

func TestBroker_Ping(t *testing.T) {
	b := Broker{
		Client: redis.NewClient(&redis.Options{
			Addr: "localhost:6379",
			DB:   0, // use default DB
		}),
	}
	if err := b.Ping(); err != nil {
		fmt.Println(err)
	}
}

func TestBroker_PushTask(t *testing.T) {

	ta := task.Task{
		Id: "id",
	}
	bytes, _ := json.Marshal(ta)
	b := Broker{
		Client: redis.NewClient(&redis.Options{
			Addr: "localhost:6379",
			DB:   0,
		}),
	}
	if err := b.PushTask(string(bytes)); err != nil {
		log.Fatal(err)
	}
	msg, err := b.PopTask()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msg)
}
