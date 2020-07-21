/*
Copyright © 2020 Ben Shi

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package out

import (
	"os"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"

	"github.com/hbish/smex/pkg/xml"
)

func TestNewWriter(t *testing.T) {
	var fs = afero.NewMemMapFs()
	w := NewWriter(fs, os.Stdout)
	assert.NotNil(t, w)
	assert.Equal(t, w.Format, Stdout)

	urls := []xml.URL{
		{Loc: "http://www.example.com/", LastMod: "2005-01-01", ChangeFreq: "monthly", Priority: 0.8},
	}
	_ = w.Write(urls, false, "test-file")
	dir, _ := afero.ReadDir(fs, "")
	assert.Equal(t, len(dir), 0)
}

func TestNewMultiWriter_NoFormats(t *testing.T) {
	var fs = afero.NewMemMapFs()
	w := NewMultiWriter(fs, os.Stdout, "")
	assert.NotNil(t, w)
	assert.Equal(t, w.Format, Stdout)

	urls := []xml.URL{
		{Loc: "http://www.example.com/", LastMod: "2005-01-01", ChangeFreq: "monthly", Priority: 0.8},
	}

	_ = w.Write(urls, false, "test-file")
	dir, _ := afero.ReadDir(fs, "")
	assert.Equal(t, len(dir), 0)
}

func TestNewMultiWriter_Csv(t *testing.T) {
	var fs = afero.NewMemMapFs()
	w := NewMultiWriter(fs, os.Stdout, "csv")
	assert.NotNil(t, w)
	assert.Equal(t, w.Format, Csv)

	urls := []xml.URL{
		{Loc: "http://www.example.com/", LastMod: "2005-01-01", ChangeFreq: "monthly", Priority: 0.8},
	}

	_ = w.Write(urls, false, "test-file")
	dir, _ := afero.ReadDir(fs, "")
	if assert.Equal(t, len(dir), 1) {
		assert.Equal(t, dir[0].Name(), "test-file.csv")
	}
}

func TestNewMultiWriter_Json(t *testing.T) {
	var fs = afero.NewMemMapFs()
	w := NewMultiWriter(fs, os.Stdout, "json")
	assert.NotNil(t, w)
	assert.Equal(t, w.Format, JSON)

	urls := []xml.URL{
		{Loc: "http://www.example.com/", LastMod: "2005-01-01", ChangeFreq: "monthly", Priority: 0.8},
	}

	_ = w.Write(urls, false, "test-file")
	dir, _ := afero.ReadDir(fs, "")
	assert.Equal(t, len(dir), 1)
	if assert.Equal(t, len(dir), 1) {
		assert.Equal(t, dir[0].Name(), "test-file.json")
	}
}
