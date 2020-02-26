package model

import (
    "fmt"
    "sync"
)

type Model struct {
    Name string `json:"name"`
}

type Cache struct {
    items map[string]interface{}
    mutex sync.Mutex
}

func NewCache() *Cache {
    return &Cache{items: map[string]interface{}{}}
}

func (c Cache) Get(key string) (interface{}, error) {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    p, ok := c.items[key]
    if !ok {
        return nil, fmt.Errorf("can not find key:%s", key)
    }
    return p, nil
}

func (c Cache) Update(key string, item interface{}) error {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    _, ok := c.items[key]
    if !ok {
        return fmt.Errorf("can not find key:%s", key)
    }
    c.items[key] = item
    return nil
}

func (c Cache) Create(key string, item interface{}) error {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    _, ok := c.items[key]
    if ok {
        return fmt.Errorf("key:%s already in cache", key)
    }
    c.items[key] = item
    return nil
}

func (c Cache) CreateOrUpdate(key string, item interface{}) (interface{}, bool) {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    if _, ok := c.items[key]; ok {
        c.items[key] = item
        return item, false
    }
    c.items[key] = item
    return item, true
}

func (c Cache) List() []interface{} {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    items := make([]interface{}, 0)
    for _, value := range c.items {
        items = append(items, value)
    }
    return items
}

func (c Cache) Delete(id string) error {
    _, ok := c.items[id]
    if ok {
        return fmt.Errorf("can not find key:%s", id)
    }
    delete(c.items, id)
    return nil
}
