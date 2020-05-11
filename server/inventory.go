package server

import (
	uuid "github.com/satori/go.uuid"
	"kobe/api"
	"sync"
)

func NewInventoryCache() *inventoryCache {
	return &inventoryCache{
		Data: map[string]*api.Inventory{},
	}
}

type inventoryCache struct {
	Data  map[string]*api.Inventory
	mutex sync.Mutex
}

func (i *inventoryCache) Put(value *api.Inventory) string {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	key := uuid.NewV4().String()
	i.Data[key] = value
	return key
}

func (i *inventoryCache) Get(id string) *api.Inventory {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	return i.Data[id]
}

func (i *inventoryCache) Delete(id string) {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	delete(i.Data, id)
}
