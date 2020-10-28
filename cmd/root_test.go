package cmd

import (
	"bytes"
	"errors"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func Test_Root_Run(t *testing.T) {
	buf := &bytes.Buffer{}
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	rootCmd.RunE = func(_ *cobra.Command, _ []string) error {
		return errors.New("")
	}

	assert.NotNil(t, Run())

	assert.Nil(t, rootRunE(rootCmd, nil))
}
