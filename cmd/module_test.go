package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Module_Run(t *testing.T) {
	at := assert.New(t)

	t.Run("success", func(t *testing.T) {
		defer func() {
			at.Nil(os.RemoveAll("testcase"))
		}()

		out, err := runCobraCmd(moduleCmd, "testcase")

		at.Nil(err)
		at.Contains(out, "Done")
	})

	t.Run("invalid module name", func(t *testing.T) {
		out, err := runCobraCmd(moduleCmd, ".")

		at.NotNil(err)
		at.Contains(out, ".")
	})
}

func Test_Module_CreateModule(t *testing.T) {
	t.Parallel()

	at := assert.New(t)

	dir, err := ioutil.TempDir("", "test_create_module")
	at.Nil(err)
	defer func() { _ = os.RemoveAll(dir) }()

	modulePath := fmt.Sprintf("%s%cmodule", dir, os.PathSeparator)

	at.NotNil(createModule(modulePath, "invalid-name/"))
}
