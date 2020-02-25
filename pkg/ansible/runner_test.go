package ansible

import (
    "fmt"
    "testing"
)

func TestPlaybookRunner_Run(t *testing.T) {
    //play := BasePlaybook{
    //    Path: "data/playbooks/test.yml",
    //}

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
    bi := NewBaseInventory(group.Hosts, []BaseGroup{group})
    fmt.Println(bi.String())

    //log.Println(inv.Json())
    //back := CallBack{
    //    onSuccess: func() {
    //        fmt.Println("success")
    //    },
    //    onError: func() {
    //        fmt.Println("error")
    //    },
    //}
    //rn := PlaybookRunner{
    //   Inventory: inv,
    //   Playbook:  play,
    //   Options:   map[string]interface{}{},
    //}
    //err := rn.Run(&back)
    //if err != nil {
    //   fmt.Println(err)
    //}

}
