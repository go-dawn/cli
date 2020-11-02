package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"time"
)

var (
	homeDir string

	execLookPath = exec.LookPath
	execCommand  = exec.Command
	osExit       = os.Exit

	skipSpinner bool
)

func init() {
	homeDir, _ = os.UserHomeDir()
}

func createFile(filePath, content string) (err error) {
	var f *os.File
	if f, err = os.Create(filePath); err != nil {
		return
	}

	defer func() { _ = f.Close() }()

	_, err = f.WriteString(content)

	return
}

func runCmd(cmd *exec.Cmd) (err error) {
	if verbose {
		var (
			stderr io.ReadCloser
			stdout io.ReadCloser
		)

		if stderr, err = cmd.StderrPipe(); err != nil {
			return
		}
		defer func() {
			_ = stderr.Close()
		}()
		go func() { _, _ = io.Copy(os.Stderr, stderr) }()

		if stdout, err = cmd.StdoutPipe(); err != nil {
			return
		}
		defer func() {
			_ = stdout.Close()
		}()
		go func() { _, _ = io.Copy(os.Stdout, stdout) }()
	}

	if err = cmd.Run(); err != nil {
		err = fmt.Errorf("failed to run %s", cmd.String())
	}

	return
}

func formatLatency(d time.Duration) time.Duration {
	switch {
	case d > time.Second:
		return d.Truncate(time.Second / 100)
	case d > time.Millisecond:
		return d.Truncate(time.Millisecond / 100)
	case d > time.Microsecond:
		return d.Truncate(time.Microsecond / 100)
	default:
		return d
	}
}

func loadConfig() (err error) {
	configFilePath := configFilePath()

	if fileExist(configFilePath) {
		if err = loadJson(configFilePath, &rc); err != nil {
			return
		}
	}

	return
}

func storeConfig() {
	_ = storeJson(configFilePath(), rc)
}

func configFilePath() string {
	if homeDir == "" {
		return configName
	}

	return fmt.Sprintf("%s%c%s", homeDir, os.PathSeparator, configName)
}

var fileExist = func(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

func storeJson(filename string, v interface{}) error {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, b, 0600)
}

func loadJson(filename string, v interface{}) error {
	b, err := ioutil.ReadFile(path.Clean(filename))
	if err != nil {
		return err
	}

	return json.Unmarshal(b, v)
}
