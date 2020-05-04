package cmd

import (
	"bytes"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestRootCmd(t *testing.T) {
	Execute()

	output, err2 := executeCommand(rootCmd, nil)
	assert.NoError(t, err2)
	assert.Contains(t, output, "smex [command]")
}

func executeCommand(root *cobra.Command, commands []string) (output string, err error) {
	buf := new(bytes.Buffer)
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs(commands)

	_, err = root.ExecuteC()

	return buf.String(), err
}
