package out

import (
	"testing"

	"github.com/hbish/smex/pkg/xml"
)

func TestSmexWriter(t *testing.T) {
	w := NewWriter()

	urls := []xml.URL{
		{Loc: "http://www.example.com/", LastMod: "2005-01-01", ChangeFreq: "monthly", Priority: 0.8},
	}

	w.Write(urls, false)
}
