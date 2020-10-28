package cmd

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_New_Run(t *testing.T) {
	at := assert.New(t)

	t.Run("new project", func(t *testing.T) {
		defer func() {
			at.Nil(os.Chdir("../"))
			at.Nil(os.RemoveAll("testcase"))
		}()

		setupCmd()
		defer teardownCmd()

		out, err := runCobraCmd(newCmd, "testcase")

		at.Nil(err)
		at.Contains(out, "Done")
	})

	t.Run("custom mod name", func(t *testing.T) {
		defer func() {
			at.Nil(os.Chdir("../"))
			at.Nil(os.RemoveAll("testcase"))
		}()

		setupCmd()
		defer teardownCmd()

		out, err := runCobraCmd(newCmd, "testcase", "custom")

		at.Nil(err)
		at.Contains(out, "custom")
	})

	t.Run("use --app and fail", func(t *testing.T) {
		defer func() {
			at.Nil(os.Chdir("../"))
			at.Nil(os.RemoveAll("testcase"))
		}()

		setupCmd(errFlag)
		defer teardownCmd()

		out, err := runCobraCmd(newCmd, "testcase", "--app")

		at.NotNil(err)
		at.Contains(out, "failed to run")
	})

	t.Run("invalid project name", func(t *testing.T) {
		out, err := runCobraCmd(newCmd, ".")

		at.NotNil(err)
		at.Contains(out, ".")
	})
}

func Test_New_CreateConfigs(t *testing.T) {
	assert.NotNil(t, createConfigs(" "))
}

func Test_New_InitGit(t *testing.T) {
	at := assert.New(t)

	t.Run("look path error", func(t *testing.T) {
		setupLookPath(errFlag)
		defer teardownLookPath()

		at.Nil(initGit(""))
	})

	t.Run("failed to create .gitignore", func(t *testing.T) {
		setupLookPath()
		defer teardownLookPath()

		at.NotNil(initGit(" "))
	})

	t.Run("failed to run command", func(t *testing.T) {
		projectPath, err := ioutil.TempDir("", "test_new_init_git")
		at.Nil(err)
		defer func() {
			at.Nil(os.RemoveAll(projectPath))
		}()

		setupLookPath()
		defer teardownLookPath()

		setupCmd(errFlag)
		defer teardownCmd()

		at.NotNil(initGit(projectPath))
	})
}
