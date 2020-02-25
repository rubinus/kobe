package ansible

type CallBackFunc func()

type CallBack struct {
    onSuccess CallBackFunc
    onError   CallBackFunc
}
