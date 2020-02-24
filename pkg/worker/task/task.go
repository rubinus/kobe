package task

type Runnable interface {
    Start() error
}

type Task struct {
    Id       string
    Runnable Runnable
}
