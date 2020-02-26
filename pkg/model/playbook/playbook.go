package playbook

import (
    "io/ioutil"
    "kobe/pkg/ansible"
    "kobe/pkg/model"
    "log"
    "path"
)

var Cache *model.Cache

func init() {
    Cache = model.NewCache()
    lookup()
}

type Playbook struct {
    *ansible.BasePlaybook
    *model.Model
}

func NewPlaybook(name string, path string) *Playbook {
    return &Playbook{
        BasePlaybook: &ansible.BasePlaybook{Path: path},
        Model:        &model.Model{Name: name},
    }
}

func lookup() {
    baseDir := "data/playbooks"
    list, err := ioutil.ReadDir(baseDir)
    log.Println("lookup playbooks")
    if err != nil {
        log.Printf("look up playbook error: %s", err)
        return
    }
    for _, item := range list {
        if !item.IsDir() {
            continue
        }
        p := path.Join(baseDir, item.Name())
        book := NewPlaybook(item.Name(), p)
        Cache.CreateOrUpdate(book.Name, book)
    }
}
