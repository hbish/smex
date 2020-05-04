package json

import (
	"os"
	"testing"

	"github.com/hbish/smex/pkg/xml"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

var urls = []xml.URL{
	{Loc: "http://www.example.com/", LastMod: "2005-01-01", ChangeFreq: "monthly", Priority: 0.8},
	{Loc: "http://www.example.com/catalog?item=12&desc=vacation_hawaii", LastMod: "", ChangeFreq: "weekly", Priority: 0},
}

var fs = afero.NewMemMapFs()

func TestNewWriter(t *testing.T) {
	// given
	actual, _ := fs.Create("test.json")
	writer := NewWriter(actual, true)

	// when
	// then
	assert.NotNil(t, writer)
}

func TestNewWriter_WriteToFile_All(t *testing.T) {
	// given
	outFile, _ := fs.Create(t.Name())
	defer outFile.Close()
	writer := NewWriter(outFile, false)

	// when
	_, _ = writer.WriteToFile(urls, false)
	actual, _ := afero.ReadFile(fs, t.Name())

	// then
	expected := "[\n{\n\"loc\": \"http://www.example.com/\",\n\"lastmod\": \"2005-01-01\",\n\"changefreq\": \"monthly\",\n\"priority\": 0.8\n},\n{\n\"loc\": \"http://www.example.com/catalog?item=12\\u0026desc=vacation_hawaii\",\n\"changefreq\": \"weekly\"\n}\n]"
	assert.Equal(t, expected, string(actual))
}

func TestNewWriter_WriteToFile_AllWithPrettyPrint(t *testing.T) {
	// given
	outFile, _ := fs.Create(t.Name())
	defer outFile.Close()
	writer := NewWriter(outFile, true)

	// when
	_, _ = writer.WriteToFile(urls, false)
	actual, _ := afero.ReadFile(fs, t.Name())

	// then
	expected := "[\n  {\n    \"loc\": \"http://www.example.com/\",\n    \"lastmod\": \"2005-01-01\",\n    \"changefreq\": \"monthly\",\n    \"priority\": 0.8\n  },\n  {\n    \"loc\": \"http://www.example.com/catalog?item=12\\u0026desc=vacation_hawaii\",\n    \"changefreq\": \"weekly\"\n  }\n]"
	assert.Equal(t, expected, string(actual))
}

func TestNewWriter_WriteToFile_OnlyLoc(t *testing.T) {
	// given
	outFile, _ := fs.Create(t.Name())
	defer outFile.Close()
	writer := NewWriter(outFile, false)

	// when
	_, _ = writer.WriteToFile(urls, true)
	actual, _ := afero.ReadFile(fs, t.Name())

	// then
	expected := "[\n{\n\"loc\": \"http://www.example.com/\"\n},\n{\n\"loc\": \"http://www.example.com/catalog?item=12\\u0026desc=vacation_hawaii\"\n}\n]"
	assert.Equal(t, expected, string(actual))
}

func TestNewWriter_WriteToFile_Error(t *testing.T) {
	// given
	_ = afero.WriteFile(fs, t.Name(), []byte("test"), os.FileMode(0400))
	file, _ := fs.OpenFile(t.Name(), 0, os.FileMode(0400))
	writer := NewWriter(file, false)

	// when
	_, err := writer.WriteToFile(urls, true)

	// then
	assert.EqualError(t, err, "write TestNewWriter_WriteToFile_Error: file handle is read only")
}
