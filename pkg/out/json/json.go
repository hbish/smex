//Package json json
package json

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
	"encoding/json"
	"io"

	"github.com/hbish/smex/pkg/xml"
)

// Writer writer
type Writer struct {
	w      io.Writer
	indent string
}

// NewWriter new json writer
func NewWriter(w io.Writer, pretty bool) *Writer {
	indent := ""
	if pretty {
		indent = "  "
	}

	return &Writer{
		w:      w,
		indent: indent,
	}
}

// WriteToFile WriteToFile
func (w Writer) WriteToFile(urls []xml.URL, loc bool) ([]byte, error) {
	if loc {
		for i, url := range urls {
			urls[i] = xml.URL{
				Loc: url.Loc,
			}
		}
	}
	marshal, err := json.MarshalIndent(urls, "", w.indent)
	if err != nil {
		return nil, err
	}
	_, err = w.w.Write(marshal)
	if err != nil {
		return nil, err
	}
	return marshal, nil
}
