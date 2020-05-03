package stdout

import (
	"fmt"
	"strings"

	"github.com/hbish/smex/pkg/xml"
)

// TODO: add ability to change delimiter
type Writer struct {
}

func NewWriter() *Writer {
	return &Writer{}
}

func (w Writer) Write(urls []xml.URL, loc bool) string {
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
	fmt.Println(sb.String())
	return sb.String()
}
