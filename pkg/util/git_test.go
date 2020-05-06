package util

import "testing"

func TestCloneRepository(t *testing.T) {
	url := "https://github.com/KubeOperator/demo.git"
	target := "tmp/"
	err := CloneRepository(url, target)
	if err != nil {
		t.Error(err)
	}
}
