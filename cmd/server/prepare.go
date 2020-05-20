package main

import (
	"github.com/KubeOperator/kobe/pkg/constant"
	"os"
	"os/exec"
)

func prepareStart() error {
	funcs := []func() error{
		makeDataDir,
		lookUpAnsibleBinPath,
		lookUpKobeInventoryBinPath,
		cleanWorkPath,
	}
	for _, f := range funcs {
		if err := f(); err != nil {
			return err
		}
	}
	return nil
}

func makeDataDir() error {
	err := os.MkdirAll(constant.ProjectDir, 0755)
	if err != nil {
		return err
	}
	return nil
}

func lookUpAnsibleBinPath() error {
	_, err := exec.LookPath(constant.AnsiblePlaybookBinPath)
	if err != nil {
		return err
	}
	return nil
}

func lookUpKobeInventoryBinPath() error {
	_, err := exec.LookPath(constant.InventoryProviderBinPath)
	if err != nil {
		return err
	}
	return nil
}

func cleanWorkPath() error {
	_ = os.RemoveAll(constant.WorkDir)
	return nil
}
