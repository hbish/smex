// Package csv csv
package csv

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

import (
	"encoding/csv"
	"fmt"
	"io"

	"github.com/hbish/smex/pkg/xml"
	"github.com/pkg/errors"
)

// Writer writer
type Writer struct {
	*csv.Writer
}

// NewWriter create a new csv writer
func NewWriter(w io.Writer, delim rune) *Writer {
	writer := csv.NewWriter(w)
	if delim != 0 {
		writer.Comma = delim
	}

	return &Writer{Writer: writer}
}

// WriteToFile WriteToFile
func (w Writer) WriteToFile(urls []xml.URL, loc bool) ([][]string, error) {
	var content [][]string
	header := w.writeHeader(loc)
	content = append(content, header)

	for _, url := range urls {
		line := []string{url.Loc}
		if !loc {
			line = append(line, url.LastMod, url.ChangeFreq, fmt.Sprintf("%.2f", url.Priority))
		}
		content = append(content, line)
	}
	err := w.WriteAll(content)
	if err != nil {
		return nil, errors.Wrap(err, "unable to write csv content")
	}
	return content, nil
}

// TODO rather than hard code fields, can use reflection or read in user input
func (w Writer) writeHeader(loc bool) []string {
	header := []string{"loc"}
	if !loc {
		header = append(header, "lastmod", "changefreq", "priority")
	}
	return header
}
