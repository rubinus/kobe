package model

import "reflect"

type Cache interface {
	List() interface{}
	Get(id string) (interface{}, error)
	Delete(id string) error
	Create(item interface{}) (interface{}, error)
	Update(item interface{}) (interface{}, error)
}

type Manager struct {
	cache Cache
	t     reflect.Type
}
