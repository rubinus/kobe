package worker

import "kobe/pkg/worker/task"

type Worker struct {
    Id     string
    config Config
}

func (w *Worker) Handler(t task.Task) error {
    t.Runnable.Start()
    return nil
}
