package inventory

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func TestNewGroup(t *testing.T) {
	group := Group{
		Name:     "test",
		Children: []Children{"a"},
		Vars: map[string]interface{}{
			"install": true,
		},
	}
	b, err := json.Marshal(group)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}
