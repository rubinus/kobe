package main

import (
	"kobe/pkg/constant"
	"os"
	"os/exec"
)

func prepareStart() error {
	funcs := []func() error{
		makeDataDir,
		lookUpAnsibleBinPath,
		lookUpKobeInventoryBinPath}
	for _, f := range funcs {
		err := f()
		return err
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
