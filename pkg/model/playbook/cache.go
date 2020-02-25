package playbook
//
//import (
//	"fmt"
//	"sync"
//)
//
//type Cache struct {
//	items map[string]interface{}
//	mutex sync.Mutex
//}
//
//func (c Cache) Get(id string) (interface{}, error) {
//	c.mutex.Lock()
//	defer c.mutex.Unlock()
//	p, ok := c.items[id]
//	if !ok {
//		return nil, fmt.Errorf("can not find key:%s", id)
//	}
//	return p, nil
//}
//
//func (c Cache) Update(item interface{}) error {
//	c.mutex.Lock()
//	defer c.mutex.Unlock()
//	_, ok := c.items[item.Id]
//	if !ok {
//		return fmt.Errorf("can not find key:%s", item.Id)
//	}
//	c.items[item.Id] = item
//	return nil
//}
//
//func (c Cache) Create(item interface{}) error {
//	c.mutex.Lock()
//	defer c.mutex.Unlock()
//	_, ok := c.items[item.Id]
//	if ok {
//		return fmt.Errorf("key:%s already in cache", item.Id)
//	}
//	c.items[item.Id] = item
//	return nil
//}
//
//func (c Cache) List() []interface{} {
//	c.mutex.Lock()
//	defer c.mutex.Unlock()
//	items := make([]interface{}, 10)
//	for _, value := range c.items {
//		items = append(items, value)
//	}
//	return items
//}
//
//func (c Cache) Delete(id string) error {
//	_, ok := c.items[id]
//	if ok {
//		return fmt.Errorf("can not find key:%s", id)
//	}
//	delete(c.items, id)
//	return nil
//}
