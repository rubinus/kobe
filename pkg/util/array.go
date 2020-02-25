package util

import (
    "reflect"
)

func Contains(obj interface{}, container interface{}) bool {
    targetValue := reflect.ValueOf(container)
    switch reflect.TypeOf(container).Kind() {
    case reflect.Slice, reflect.Array:
        for i := 0; i < targetValue.Len(); i++ {
            if targetValue.Index(i).Interface() == obj {
                return true
            }
        }
    case reflect.Map:
        if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
            return true
        }
    }
    return false
}
