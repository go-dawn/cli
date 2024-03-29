package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

var (
	needError bool
	errFlag   = struct{}{}
)

func init() {
	verbose = true
}

func fakeExecCommand(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestHelperProcess", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1"}
	if needError {
		cmd.Env = append(cmd.Env, "GO_WANT_HELPER_NEED_ERR=1")
	}
	return cmd
}

func TestHelperProcess(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}
	args := os.Args
	for len(args) > 0 {
		if args[0] == "--" {
			args = args[1:]
			break
		}
		args = args[1:]
	}

	if len(args) == 0 {
		_, _ = fmt.Fprintf(os.Stderr, "No command")
		os.Exit(2)
	}

	if os.Getenv("GO_WANT_HELPER_NEED_ERR") == "1" {
		_, _ = fmt.Fprintf(os.Stderr, "fake error")
		os.Exit(1)
	}

	os.Exit(0)
}

func setupCmd(flag ...struct{}) {
	execCommand = fakeExecCommand
	if len(flag) > 0 {
		needError = true
	}
}

func teardownCmd() {
	execCommand = exec.Command
	needError = false
}

func setupLookPath(flag ...struct{}) {
	execLookPath = func(file string) (s string, err error) {
		if len(flag) > 0 {
			err = errors.New("fake look path error")
		}
		return
	}
}

func teardownLookPath() {
	execLookPath = exec.LookPath
}

func runCobraCmd(cmd *cobra.Command, args ...string) (string, error) {
	b := new(bytes.Buffer)

	cmd.ResetCommands()
	cmd.SetErr(b)
	cmd.SetOut(b)
	cmd.SetArgs(args)
	err := cmd.Execute()

	return b.String(), err
}

func setupHomeDir(t *testing.T, pattern string) string {
	homeDir, err := ioutil.TempDir("", "test_"+pattern)
	assert.Nil(t, err)
	return homeDir
}

func teardownHomeDir(dir string) {
	_ = os.RemoveAll(dir)
}

func setupSpinner() {
	skipSpinner = true
}

func teardownSpinner() {
	skipSpinner = false
}
