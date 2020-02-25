package inventory

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func changeName(b *[5]string) {
	b[2] = "2"
}

func TestNewInventory(t *testing.T) {
	host := Host{
		Name:     "test",
		Ip:       "172.16.10.142",
		Port:     22,
		Username: "root",
		Password: "Calong@2015",
		Vars: map[string]interface{}{
			"connection": "local",
		},
	}
	group := Group{
		Name:     "test",
		Children: []Children{"host"},
		Vars: map[string]interface{}{
			"install": true,
		},
	}
	hosts := []Host{host}
	groups := []Group{group}
	inventory := NewInventory(groups, hosts)
	b, err := json.Marshal(inventory)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}
