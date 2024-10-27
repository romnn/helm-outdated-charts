package cmd

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
)

var errCmdNotInstalled = errors.New("command not installed")

// Command ...
type Command struct {
	cmd         string
	defaultArgs []string
}

// New returns the Command or an error.
func New(cmd string, defaultArgs ...string) (*Command, error) {
	c := &Command{
		cmd:         cmd,
		defaultArgs: defaultArgs,
	}

	return c, c.verify()
}

// Run starts the command, waits until it finished and returns stdOut or an error containing the stdError message.
func (c *Command) Run(args ...string) (string, error) {
	cmd := exec.Command(c.cmd, append(c.defaultArgs, args...)...)

	if v, ok := os.LookupEnv("DEBUG"); ok && v == "true" {
		fmt.Println("running: ", cmd.String())
	}

	var (
		stdOut,
		stdErr bytes.Buffer
	)
	cmd.Stdout = &stdOut
	cmd.Stderr = &stdErr

	if err := cmd.Start(); err != nil {
		return "", errors.Wrap(err, strings.TrimSpace(stdErr.String()))
	}

	if err := cmd.Wait(); err != nil {
		return "", errors.Wrap(err, strings.TrimSpace(stdErr.String()))
	}
	return strings.TrimSpace(stdOut.String()), nil
}

// verify checks if the command is installed.
func (c *Command) verify() error {
	res, err := c.Run("version")
	if err != nil && strings.Contains(err.Error(), "not found") ||
		strings.Contains(res, "not found") {
		return errCmdNotInstalled
	}

	return nil
}
