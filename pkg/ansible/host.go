package ansible

type BaseHost struct {
    Hostname string
    Vars     map[string]interface{}
}

func (bh BaseHost) Data() map[string]interface{} {
    hostData := make(map[string]interface{})
    hostData[bh.Hostname] = bh.Vars
    return hostData
}
