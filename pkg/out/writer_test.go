package out

import (
	"bytes"
	"testing"
)

func TestSmexWriter(t *testing.T) {
	b := &bytes.Buffer{}
	w := NewWriter(b)
	w.Write("123456")
}
