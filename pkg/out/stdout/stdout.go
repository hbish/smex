// Package stdout stdout
package stdout

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
	"fmt"
	"io"
	"strings"

	"github.com/hbish/smex/pkg/xml"
)

// Writer writer
// TODO: add ability to change delimiter
type Writer struct {
	w io.Writer
}

// NewWriter new stdout writer
func NewWriter(w io.Writer) *Writer {
	return &Writer{
		w: w,
	}
}

// Write Write
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

	for _, url := range urls {
		sb.WriteString(fmt.Sprintf("%-*s\t", maxLocLength, url.Loc))
		if loc {
			sb.WriteString(fmt.Sprintf("\n"))
			continue
		}
		if url.LastMod != "" {
			sb.WriteString(fmt.Sprintf("%-20s", url.LastMod))
		}
		if url.ChangeFreq != "" {
			sb.WriteString(fmt.Sprintf("%-7s\t", url.ChangeFreq))
		}
		if url.Priority != 0 {
			sb.WriteString(fmt.Sprintf("%.2f", url.Priority))
		}
		sb.WriteString(fmt.Sprintf("\n"))

		if len(url.News) > 0 {
			writeNewsLine(&sb, url.News)
		}

		if len(url.Image) > 0 {
			writeImageLine(&sb, url.Image)
		}

		if len(url.Video) > 0 {
			writeVideoLine(&sb, url.Video)
		}
	}
	_, err := w.w.Write([]byte(sb.String()))
	return err
}

func writeImageLine(sb *strings.Builder, images []xml.Image) {
	for _, image := range images {
		sb.WriteString(fmt.Sprintf("Image\t"))
		if image.Title != "" {
			sb.WriteString(fmt.Sprintf("%s\t", image.Title))
		}
		if image.Caption != "" {
			sb.WriteString(fmt.Sprintf("%s\t", image.Caption))
		}
		sb.WriteString(fmt.Sprintf("%s\n", image.Loc))
	}
}

func writeVideoLine(sb *strings.Builder, videos []xml.Video) {
	for _, video := range videos {
		sb.WriteString(fmt.Sprintf("Video\t%s\n", video.ContentLoc))
	}
}

func writeNewsLine(sb *strings.Builder, news []xml.News) {
	for _, n := range news {
		sb.WriteString(fmt.Sprintf("News\t%s\n", n.Title))
	}
}
