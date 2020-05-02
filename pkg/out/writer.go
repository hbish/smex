package out

import (
	"bufio"
	"io"
)

type outType string

const (
	Stdout outType = "stdout"
	Csv    outType = "csv"
	Json   outType = "json"
)

// A SmexWriter writes records to different formats.
//
// Returns a particular writer based on the type
type SmexWriter struct {
	Type string
	w    *bufio.Writer
}

func (w SmexWriter) Write(s string) {

}

// NewWriter returns a new SmexWriter that writes to w.
func NewWriter(w io.Writer) *SmexWriter {
	return &SmexWriter{
		Type: "csv",
		w:    bufio.NewWriter(w),
	}
}
