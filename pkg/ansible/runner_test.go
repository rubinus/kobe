package ansible

import (
    "log"
    "testing"
)

func TestPlaybookRunner_Run(t *testing.T) {
    play := BasePlaybook{
        Path: "/Users/shenchenyang/go/src/github.com/kobe/data/playbooks/test/test.yml",
    }

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
    c := CallBack{
        onError: func() {
            log.Println("error")
        },
        onSuccess: func() {
            log.Println("success")
        },
    }
    bi := NewBaseInventory(group.Hosts, []BaseGroup{group})
    runner := PlaybookRunner{
        BasePlaybook:  &play,
        BaseInventory: bi,
        WorkPath:      "/Users/shenchenyang/go/src/github.com/kobe/data/task/1/",
        Options:       map[string]interface{}{},
    }
    _ = runner.Run(&c)
}
