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
	"bytes"
	"testing"

	"github.com/hbish/smex/pkg/xml"
	"github.com/stretchr/testify/assert"
)

var urls = []xml.URL{
	{Loc: "http://www.example.com/", LastMod: "2005-01-01", ChangeFreq: "monthly", Priority: 0.8},
	{Loc: "http://www.example.com/catalog?item=12&desc=vacation_hawaii", LastMod: "", ChangeFreq: "weekly", Priority: 0},
	{Loc: "http://www.example.com/catalog?item=73&desc=vacation_new_zealand", LastMod: "2004-12-23", ChangeFreq: "weekly", Priority: 0},
	{Loc: "http://www.example.com/catalog?item=74&desc=vacation_newfoundland", LastMod: "2004-12-23T18:00:15+00:00", ChangeFreq: "", Priority: 0.3},
	{Loc: "http://www.example.com/catalog?item=83&desc=vacation_usa", LastMod: "2004-11-23", ChangeFreq: "", Priority: 0},
}

var urlsComplex = []xml.URL{
	{Loc: "http://www.example.com/", LastMod: "2005-01-01", ChangeFreq: "monthly", Priority: 0.8,
		Image: []xml.Image{{Loc: "http://www.example.com/example.jpg", Title: "", Caption: "", GeoLocation: "", License: ""}},
		Video: []xml.Video{{ThumbnailLoc: "", Title: "example title", Description: "example description", ContentLoc: "http://www.example.com/v/video", PlayerLoc: "http://www.example.com/v/video", Duration: 110, ExpirationDate: "", Rating: 0, ViewCount: 0, PublicationDate: "2010-10-05T18:52:47.000Z", FamilyFriendly: "", Restriction: "", Platform: "", Price: 0, RequiresSubscription: "", Uploader: "", Live: "", Tag: []string(nil), Category: ""}},
		News:  []xml.News{{Publication: xml.Publication{Name: "The Example", Language: "en"}, PublicationDate: "2020-05-05T13:41:54Z", Title: "The Great Example", Genres: "", Keywords: "", Access: "Subscription", StockTickers: ""}}},
}

func TestNewWriter(t *testing.T) {
	//given
	buf := new(bytes.Buffer)
	writer := NewWriter(buf)
	//when
	//then
	assert.NotNil(t, writer)
}

func TestWriter_Write_All_Simple(t *testing.T) {
	// given
	buf := new(bytes.Buffer)
	writer := NewWriter(buf)

	// when
	_ = writer.Write(urls, false)

	// then
	expected := "http://www.example.com/                                          \t2005-01-01          monthly\t0.80\n" +
		"http://www.example.com/catalog?item=12&desc=vacation_hawaii      \tweekly \t\n" +
		"http://www.example.com/catalog?item=73&desc=vacation_new_zealand \t2004-12-23          weekly \t\n" +
		"http://www.example.com/catalog?item=74&desc=vacation_newfoundland\t2004-12-23T18:00:15+00:000.30\n" +
		"http://www.example.com/catalog?item=83&desc=vacation_usa         \t2004-11-23          \n"
	assert.Equal(t, expected, buf.String())
}

func TestWriter_Write_All_Complex(t *testing.T) {
	// given
	buf := new(bytes.Buffer)
	writer := NewWriter(buf)

	// when
	_ = writer.Write(urlsComplex, false)

	// then
	expected := "http://www.example.com/\t2005-01-01          monthly\t0.80\n" +
		"News\tThe Great Example\nImage\thttp://www.example.com/example.jpg\n" +
		"Video\thttp://www.example.com/v/video\n"
	assert.Equal(t, expected, buf.String())
}

func TestWriter_Write_OnlyLoc(t *testing.T) {
	// given
	buf := new(bytes.Buffer)
	writer := NewWriter(buf)

	// when
	_ = writer.Write(urls, true)

	// then
	expected := "http://www.example.com/\t\n" +
		"http://www.example.com/catalog?item=12&desc=vacation_hawaii\t\n" +
		"http://www.example.com/catalog?item=73&desc=vacation_new_zealand\t\n" +
		"http://www.example.com/catalog?item=74&desc=vacation_newfoundland\t\n" +
		"http://www.example.com/catalog?item=83&desc=vacation_usa\t\n"
	assert.Equal(t, expected, buf.String())
}
