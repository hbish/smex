package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractCmd_Error(t *testing.T) {
	_, err := executeCommand(rootCmd, []string{"extract"})
	assert.EqualError(t, err, "extract expects the location of the sitemap")
}

func TestExtractCmd(t *testing.T) {
	output, err := executeCommand(rootCmd, []string{"extract", "../testdata/all_in_one_sitemap.xml", "--loc"})
	assert.NoError(t, err)
	assert.Contains(t, output, "example.com")
}

func TestExtractCmd_Help(t *testing.T) {
	output, err := executeCommand(rootCmd, []string{"extract", "--help"})
	assert.NoError(t, err)
	assert.Contains(t, output, "smex extract [URI] [flags]")
}
