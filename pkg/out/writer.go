// Package out out
package out

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
	"io"

	"github.com/hbish/smex/pkg/out/json"
	"github.com/hbish/smex/pkg/out/stdout"

	"github.com/spf13/afero"

	"github.com/hbish/smex/pkg/out/csv"

	"github.com/hbish/smex/pkg/xml"
)

// Format for the output
type Format string

const (
	// Stdout stdout
	Stdout Format = "stdout"
	// Csv csv
	Csv Format = "csv"
	// JSON json
	JSON Format = "json"
)

// SmexWriter writes records to different formats.
// Returns a particular writer based on the type
type SmexWriter struct {
	fs     afero.Fs
	w      io.Writer
	Format Format
}

// NewWriter returns a new SmexWriter that writes to stdout
func NewWriter(fs afero.Fs, w io.Writer) *SmexWriter {
	return &SmexWriter{
		fs:     fs,
		w:      w,
		Format: Stdout,
	}
}

// NewMultiWriter returns a new SmexWriter that can be configured
// to write to multiple destinations.
func NewMultiWriter(fs afero.Fs, w io.Writer, outFormat string) *SmexWriter {
	if outFormat == "" {
		return &SmexWriter{
			fs:     fs,
			w:      w,
			Format: Stdout,
		}
	}

	return &SmexWriter{
		fs:     fs,
		w:      w,
		Format: Format(outFormat),
	}
}

// write a slice of URL
func (w SmexWriter) Write(urls []xml.URL, loc bool, filename string) error {

	switch w.Format {
	case Csv:
		csvFile, err := w.fs.Create(filename + ".csv")
		if err != nil {
			return err
		}
		defer csvFile.Close()
		writer := csv.NewWriter(csvFile, ',')
		defer writer.Flush()
		_, _ = writer.WriteToFile(urls, loc)
	case JSON:
		jsonFile, err := w.fs.Create(filename + ".json")
		if err != nil {
			return err
		}
		defer jsonFile.Close()
		writer := json.NewWriter(jsonFile, true)
		_, _ = writer.WriteToFile(urls, loc)
	default:
		writer := stdout.NewWriter(w.w)
		err := writer.Write(urls, loc)
		if err != nil {
			return err
		}
	}

	return nil
}
