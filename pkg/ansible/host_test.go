package ansible

import (
    "fmt"
    "testing"
)

func TestBaseHost_Data(t *testing.T) {
    host := BaseHost{
        Hostname: "test",
        Vars: map[string]interface{}{
            "ansible_ssh_user": "root",
            "ansible_ssh_pass": "Calong@2015",
            "ansible_ssh_port": 22,
            "ansible_ssh_host": "172.16.10.142",
        },
    }
    fmt.Print(host.Data())
}
