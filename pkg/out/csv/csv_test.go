package csv

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
	actual, _ := fs.Create("test.csv")
	writer := NewWriter(actual, 0)

	// when
	// then
	assert.NotNil(t, writer)
}

func TestNewWriter_WriteToFile_All(t *testing.T) {
	// given
	outFile, _ := fs.Create(t.Name())
	defer outFile.Close()
	writer := NewWriter(outFile, ',')

	// when
	actualContent, _ := writer.WriteToFile(urls, false)
	actual, _ := afero.ReadFile(fs, t.Name())

	// then
	assert.Equal(t, len(actualContent), 3)

	expected := "loc,lastmod,changefreq,priority\n" +
		"http://www.example.com/,2005-01-01,monthly,0.80\n" +
		"http://www.example.com/catalog?item=12&desc=vacation_hawaii,,weekly,0.00\n"
	assert.Equal(t, expected, string(actual))
}

func TestNewWriter_WriteToFile_AllWithPipeDelimiter(t *testing.T) {
	// given
	outFile, _ := fs.Create(t.Name())
	defer outFile.Close()
	writer := NewWriter(outFile, '|')

	// when
	actualContent, _ := writer.WriteToFile(urls, false)
	actual, _ := afero.ReadFile(fs, t.Name())

	// then
	assert.Equal(t, len(actualContent), 3)

	expected := "loc|lastmod|changefreq|priority\n" +
		"http://www.example.com/|2005-01-01|monthly|0.80\n" +
		"http://www.example.com/catalog?item=12&desc=vacation_hawaii||weekly|0.00\n"
	assert.Equal(t, expected, string(actual))
}

func TestNewWriter_WriteToFile_OnlyLoc(t *testing.T) {
	// given
	outFile, _ := fs.Create(t.Name())
	defer outFile.Close()
	writer := NewWriter(outFile, ',')

	// when
	actualContent, _ := writer.WriteToFile(urls, true)
	actual, _ := afero.ReadFile(fs, t.Name())

	// then
	assert.Equal(t, len(actualContent), 3)
	expected := "loc\n" +
		"http://www.example.com/\n" +
		"http://www.example.com/catalog?item=12&desc=vacation_hawaii\n"
	assert.Equal(t, expected, string(actual))
}

func TestNewWriter_WriteToFile_Error(t *testing.T) {
	// given
	_ = afero.WriteFile(fs, t.Name(), []byte("test"), os.FileMode(0400))
	file, _ := fs.OpenFile(t.Name(), 0, os.FileMode(0400))
	writer := NewWriter(file, ',')

	// when
	_, err := writer.WriteToFile(urls, true)

	// then
	assert.EqualError(t, err, "unable to write csv content: write TestNewWriter_WriteToFile_Error: file handle is read only")
}
