package ansible

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"
)

type AnsibleExecuter struct {
	Write io.Writer
}

func (e *AnsibleExecuter) Execute(command string, args []string, prefix string) error {
	if e.Write == nil {
		e.Write = os.Stdout
	}

	cmd := exec.Command(command, args...)
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		return errors.New("can not read stdout reason  " + err.Error())
	}

	cmdErrorReader, err := cmd.StderrPipe()
	if err != nil {
		return errors.New("can not read stderr reason  " + err.Error())
	}

	scanner := bufio.NewScanner(cmdReader)
	errorScanner := bufio.NewScanner(cmdErrorReader)
	go func() {
		for scanner.Scan() {
			_, _ = fmt.Fprint(e.Write, "\n", scanner.Text())
		}
		for errorScanner.Scan() {
			_, _ = fmt.Fprint(e.Write, "\n", errorScanner.Text())
		}
	}()
	timeInit := time.Now()
	err = cmd.Start()
	if err != nil {
		return errors.New("(Execute Error) -> " + err.Error())
	}

	err = cmd.Wait()
	elapsedTime := time.Since(timeInit)
	if err != nil {
		return errors.New("(Execute Error) -> " + err.Error())
	}

	_, _ = fmt.Fprintf(e.Write, "Duration: %s\n", elapsedTime.String())
	return nil
}
