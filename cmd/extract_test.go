package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractCmd_Error(t *testing.T) {
	_, err := executeCommand(rootCmd, []string{"extract"})
	assert.EqualError(t, err, "extract expects the location of the sitemap\n")
}

func TestExtractCmd(t *testing.T) {
	output, err := executeCommand(rootCmd, []string{"extract", "../testdata/yoast_post_sitemap.xml", "--loc"})
	assert.NoError(t, err)
	assert.Contains(t, output, "yoast.com")
}

func TestExtractCmd_Help(t *testing.T) {
	output, err := executeCommand(rootCmd, []string{"extract", "--help"})
	assert.NoError(t, err)
	assert.Contains(t, output, "smex extract [URI] [flags]")
}
