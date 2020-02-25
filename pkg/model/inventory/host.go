package inventory

const (
	DEFAULT_SSH_PORT = 22
)

type Host struct {
	Name     string                 `json:"name"`
	Ip       string                 `json:"host"`
	Port     int                 `json:"port"`
	Password string                 `json:"password"`
	Username string                 `json:"username"`
	Vars     map[string]interface{} `json:"vars"`
}
