package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

type commander struct{}

func (this *commander) run(commandName string, arg ...string) error {
	cmd := exec.Command(commandName, arg...)
	output, err := cmd.Output()
	if err != nil {
		return err
	}

	fmt.Println(outputAsString(output))
	return nil
}

func outputAsString(output []byte) string {
	buffer := bytes.NewBuffer(output)
	return buffer.String()
}
