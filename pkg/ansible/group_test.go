package ansible

import (
    "fmt"
    "testing"
)

func TestBaseGroup_Data(t *testing.T) {
    host := BaseHost{
        Hostname: "test",
        Vars: map[string]interface{}{
            "ansible_ssh_user": "root",
            "ansible_ssh_pass": "Calong@2015",
            "ansible_ssh_port": 22,
            "ansible_ssh_host": "172.16.10.142",
        },
    }
    group := BaseGroup{
        Name:     "test",
        Vars:     map[string]interface{}{},
        Hosts:    []BaseHost{host},
        Children: []BaseGroup{},
    }
    fmt.Println(group.Children)
    fmt.Println(group.Data())
}
