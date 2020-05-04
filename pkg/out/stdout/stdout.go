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
package stdout

import (
	"fmt"
	"io"
	"strings"

	"github.com/hbish/smex/pkg/xml"
)

// TODO: add ability to change delimiter
type Writer struct {
	w io.Writer
}

func NewWriter(w io.Writer) *Writer {
	return &Writer{
		w: w,
	}
}

func (w Writer) Write(urls []xml.URL, loc bool) error {
	var maxLocLength int
	var sb strings.Builder

	if !loc {
		for _, line := range urls {
			if len(line.Loc) > maxLocLength {
				maxLocLength = len(line.Loc)
			}
		}
	}

	for _, line := range urls {
		sb.WriteString(fmt.Sprintf("%-*s\t", maxLocLength, line.Loc))
		if loc {
			sb.WriteString(fmt.Sprintf("\n"))
			continue
		}
		if line.LastMod != "" {
			sb.WriteString(fmt.Sprintf("%-20s", line.LastMod))
		}
		if line.ChangeFreq != "" {
			sb.WriteString(fmt.Sprintf("%-7s\t", line.ChangeFreq))
		}
		if line.Priority != 0 {
			sb.WriteString(fmt.Sprintf("%.2f", line.Priority))
		}
		sb.WriteString(fmt.Sprintf("\n"))
	}
	_, err := w.w.Write([]byte(sb.String()))
	return err
}
