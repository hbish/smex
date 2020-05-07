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
package xml

import (
	"encoding/xml"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromXMLWithExampleXML(t *testing.T) {
	// sourced from https://www.sitemaps.org/protocol.html
	input := `<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
    <url>
        <loc>http://www.example.com/</loc>
        <lastmod>2005-01-01</lastmod>
        <changefreq>monthly</changefreq>
        <priority>0.8</priority>
    </url>
</urlset>
`

	actual, err := UnmarshalXMLP([]byte(input), "")

	var expected = URLSet{
		XMLName: xml.Name{Space: "http://www.sitemaps.org/schemas/sitemap/0.9", Local: "urlset"},
		XMLNs:   "http://www.sitemaps.org/schemas/sitemap/0.9",
		URL: []URL{
			{Loc: "http://www.example.com/", LastMod: "2005-01-01", ChangeFreq: "monthly", Priority: 0.8},
		},
	}
	if assert.NotNil(t, actual) && assert.Nil(t, err) {
		assert.EqualValues(t, *actual, expected)
	}
}

func TestFromXMLWithExampleXMLMultipleURLs(t *testing.T) {
	input := `<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
    <url>
        <loc>http://www.example.com/</loc>
        <lastmod>2005-01-01</lastmod>
        <changefreq>monthly</changefreq>
        <priority>0.8</priority>
    </url>
    <url>
        <loc>http://www.example.com/catalog?item=12&amp;desc=vacation_hawaii</loc>
        <changefreq>weekly</changefreq>
    </url>
    <url>
        <loc>http://www.example.com/catalog?item=73&amp;desc=vacation_new_zealand</loc>
        <lastmod>2004-12-23</lastmod>
        <changefreq>weekly</changefreq>
    </url>
    <url>
        <loc>http://www.example.com/catalog?item=74&amp;desc=vacation_newfoundland</loc>
        <lastmod>2004-12-23T18:00:15+00:00</lastmod>
        <priority>0.3</priority>
    </url>
    <url>
        <loc>http://www.example.com/catalog?item=83&amp;desc=vacation_usa</loc>
        <lastmod>2004-11-23</lastmod>
    </url>
</urlset>
`

	t.Run("No Filter", func(t *testing.T) {
		actual, err := UnmarshalXMLP([]byte(input), "")

		var expected = URLSet{
			XMLName: xml.Name{Space: "http://www.sitemaps.org/schemas/sitemap/0.9", Local: "urlset"},
			XMLNs:   "http://www.sitemaps.org/schemas/sitemap/0.9",
			URL: []URL{
				{Loc: "http://www.example.com/", LastMod: "2005-01-01", ChangeFreq: "monthly", Priority: 0.8},
				{Loc: "http://www.example.com/catalog?item=12&desc=vacation_hawaii", LastMod: "", ChangeFreq: "weekly", Priority: 0},
				{Loc: "http://www.example.com/catalog?item=73&desc=vacation_new_zealand", LastMod: "2004-12-23", ChangeFreq: "weekly", Priority: 0},
				{Loc: "http://www.example.com/catalog?item=74&desc=vacation_newfoundland", LastMod: "2004-12-23T18:00:15+00:00", ChangeFreq: "", Priority: 0.3},
				{Loc: "http://www.example.com/catalog?item=83&desc=vacation_usa", LastMod: "2004-11-23", ChangeFreq: "", Priority: 0},
			},
		}
		if assert.NotNil(t, actual) && assert.Nil(t, err) {
			assert.EqualValues(t, *actual, expected)
		}
	})

	t.Run("Working Filter", func(t *testing.T) {
		actual, err := UnmarshalXMLP([]byte(input), "catalog")

		var expected = URLSet{
			XMLName: xml.Name{Space: "http://www.sitemaps.org/schemas/sitemap/0.9", Local: "urlset"},
			XMLNs:   "http://www.sitemaps.org/schemas/sitemap/0.9",
			URL: []URL{
				{Loc: "http://www.example.com/catalog?item=12&desc=vacation_hawaii", LastMod: "", ChangeFreq: "weekly", Priority: 0},
				{Loc: "http://www.example.com/catalog?item=73&desc=vacation_new_zealand", LastMod: "2004-12-23", ChangeFreq: "weekly", Priority: 0},
				{Loc: "http://www.example.com/catalog?item=74&desc=vacation_newfoundland", LastMod: "2004-12-23T18:00:15+00:00", ChangeFreq: "", Priority: 0.3},
				{Loc: "http://www.example.com/catalog?item=83&desc=vacation_usa", LastMod: "2004-11-23", ChangeFreq: "", Priority: 0},
			},
		}
		if assert.NotNil(t, actual) && assert.Nil(t, err) {
			assert.EqualValues(t, *actual, expected)
		}
	})

	t.Run("Invalid Filter", func(t *testing.T) {
		_, err := UnmarshalXMLP([]byte(input), "*")

		assert.EqualError(t, err, "error parsing regexp: missing argument to repetition operator: `*`")
	})
}

func TestFromXMLWithAllInOneExample(t *testing.T) {
	input, _ := ioutil.ReadFile("../../testdata/all_in_one_sitemap.xml")

	actual, err := UnmarshalXMLP([]byte(input), "")

	var expected = URLSet{
		XMLName: xml.Name{Space: "http://www.sitemaps.org/schemas/sitemap/0.9", Local: "urlset"},
		XMLNs:   "http://www.sitemaps.org/schemas/sitemap/0.9",
		URL: []URL{{Loc: "http://www.example.com/", LastMod: "2020-05-05T12:41:54Z", ChangeFreq: "always", Priority: 1,
			Image: []Image{{Loc: "http://www.example.com/example.jpg", Title: "", Caption: "", GeoLocation: "", License: ""}},
			Video: []Video{{ThumbnailLoc: "", Title: "example title", Description: "example description", ContentLoc: "", PlayerLoc: "http://www.example.com/v/video", Duration: 110, ExpirationDate: "", Rating: 0, ViewCount: 0, PublicationDate: "2010-10-05T18:52:47.000Z", FamilyFriendly: "", Restriction: "", Platform: "", Price: 0, RequiresSubscription: "", Uploader: "", Live: "", Tag: []string(nil), Category: ""}},
			News:  []News{{Publication: Publication{Name: "The Example", Language: "en"}, PublicationDate: "2020-05-05T13:41:54Z", Title: "The Great Example", Genres: "", Keywords: "", Access: "Subscription", StockTickers: ""}}}},
	}

	if assert.NotNil(t, actual) && assert.Nil(t, err) {
		assert.EqualValues(t, *actual, expected)
	}
}

func TestFromXMLWithInvalidHTML(t *testing.T) {
	input := `<html></html>`

	actual, err := UnmarshalXMLP([]byte(input), "")
	assert.Nil(t, actual)
	assert.EqualError(t, err, "expected element type <urlset> but have <html>")
}

func TestFromXMLWithInvalidRandomString(t *testing.T) {
	input := `easy as 123`

	actual, err := UnmarshalXMLP([]byte(input), "")
	assert.Nil(t, actual)
	assert.EqualError(t, err, "EOF")
}
