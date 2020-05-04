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
	"github.com/hbish/smex/pkg/out/json"
	"github.com/spf13/afero"

	"github.com/hbish/smex/pkg/out/csv"

	"github.com/hbish/smex/pkg/out/stdout"
	"github.com/hbish/smex/pkg/xml"
)

type Format string

const (
	Stdout Format = "stdout"
	Csv    Format = "csv"
	Json   Format = "json"
)

// A SmexWriter writes records to different formats.
//
// Returns a particular writer based on the type
type SmexWriter struct {
	fs      afero.Fs
	Formats []Format
}

// NewWriter returns a new SmexWriter that writes to stdout
func NewWriter(fs afero.Fs) *SmexWriter {
	return &SmexWriter{
		fs:      fs,
		Formats: []Format{Stdout},
	}
}

// NewMultiWriter returns a new SmexWriter that can be configured
// to write to multiple destinations.
func NewMultiWriter(fs afero.Fs, outFormats []string) *SmexWriter {
	if len(outFormats) == 0 {
		return NewWriter(fs)
	}
	var formats []Format
	for _, t := range outFormats {
		formats = append(formats, Format(t))
	}

	return &SmexWriter{
		fs:      fs,
		Formats: formats,
	}
}

func (w SmexWriter) Write(urls []xml.URL, loc bool) error {
	if isFormatRequested(w.Formats, Stdout) {
		writer := stdout.NewWriter()
		writer.Write(urls, loc)
	}

	if isFormatRequested(w.Formats, Csv) {
		csvFile, err := w.fs.Create("smex-output.csv")
		if err != nil {
			return err
		}
		defer csvFile.Close()
		writer := csv.NewWriter(csvFile, ',')
		defer writer.Flush()
		_, _ = writer.WriteToFile(urls, loc)
	}

	if isFormatRequested(w.Formats, Json) {
		jsonFile, err := w.fs.Create("smex-output.json")
		if err != nil {
			return err
		}
		defer jsonFile.Close()
		writer := json.NewWriter(jsonFile, true)
		_, _ = writer.WriteToFile(urls, loc)
	}

	return nil
}

func isFormatRequested(fs []Format, format Format) bool {
	for _, f := range fs {
		if format == f {
			return true
		}
	}
	return false
}
