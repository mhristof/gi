package util

import (
	"bytes"
	"os/exec"

	"github.com/pkg/errors"
)

func Bash(command string) (string, string, error) {
	var stdout, stderr bytes.Buffer

	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return stdout.String(), stderr.String(), errors.Wrap(err, "cannot execute command")
	}

	return stdout.String(), stderr.String(), nil
}
