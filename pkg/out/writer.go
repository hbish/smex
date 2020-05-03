package out

import (
	"bufio"
	"fmt"

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
	Formats []Format
	w       *bufio.Writer
}

func (w SmexWriter) Write(urls []xml.URL, loc bool) {
	for _, f := range w.Formats {
		switch f {
		case Stdout:
			writer := stdout.NewWriter()
			writer.Write(urls, loc)
		case Csv:
			fmt.Println(urls)
		case Json:
			fmt.Println(urls)
		default:
		}
	}
}

// NewWriter returns a new SmexWriter that writes to w.
func NewWriter() *SmexWriter {
	return &SmexWriter{
		Formats: []Format{Stdout},
	}
}

// NewMultiWriter returns a new SmexWriter that can be configured
// to write to multiple destinations.
func NewMultiWriter(outFormats []string) *SmexWriter {
	if len(outFormats) == 0 {
		return NewWriter()
	}

	var formats []Format
	for _, t := range outFormats {
		formats = append(formats, Format(t))
	}

	return &SmexWriter{
		Formats: formats,
	}
}
