package task

type Runnable interface {
	Start() error
	Stop() error
}

type Task struct {
	Id       string
	Runnable Runnable
}
