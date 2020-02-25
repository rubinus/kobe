package playbook

import (
	"fmt"
	"kobe/pkg/ansible"
	"reflect"
	"testing"
)

func TestPlaybook(t *testing.T) {
	p := Playbook{
		Id:   "123",
		Name: "123",
		base: ansible.BasePlaybook{},
	}
	fmt.Println(reflect.TypeOf(p))
}
