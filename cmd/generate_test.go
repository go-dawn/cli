package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Generate_Run(t *testing.T) {
	out, err := runCobraCmd(generateCmd)

	assert.Nil(t, err)
	assert.Contains(t, out, "generate")
}
