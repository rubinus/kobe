package inventory

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func TestNewHost(t *testing.T) {
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
	b, err := json.Marshal(host)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}
