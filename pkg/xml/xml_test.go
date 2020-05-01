package xml

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"testing"
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

	actual, err := FromXML([]byte(input))

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

	actual, err := FromXML([]byte(input))

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
}

func TestFromXMLWithReal(t *testing.T) {
	input := `<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9" xmlns:news="http://www.google.com/schemas/sitemap-news/0.9" xmlns:xhtml="http://www.w3.org/1999/xhtml" xmlns:mobile="http://www.google.com/schemas/sitemap-mobile/1.0" xmlns:image="http://www.google.com/schemas/sitemap-image/1.1" xmlns:video="http://www.google.com/schemas/sitemap-video/1.1">
<url> <loc>https://hbish.com/about/</loc> <changefreq>daily</changefreq> <priority>0.7</priority> </url>
<url> <loc>https://hbish.com/talks/</loc> <changefreq>daily</changefreq> <priority>0.7</priority> </url>
</urlset>
`

	actual, err := FromXML([]byte(input))

	var expected = URLSet{
		XMLName: xml.Name{Space: "http://www.sitemaps.org/schemas/sitemap/0.9", Local: "urlset"},
		XMLNs:   "http://www.sitemaps.org/schemas/sitemap/0.9",
		URL: []URL{
			{Loc: "https://hbish.com/about/", LastMod: "", ChangeFreq: "daily", Priority: 0.7},
			{Loc: "https://hbish.com/talks/", LastMod: "", ChangeFreq: "daily", Priority: 0.7},
		},
	}
	if assert.NotNil(t, actual) && assert.Nil(t, err) {
		assert.EqualValues(t, *actual, expected)
	}
}

func TestFromXMLWithInvalidHTML(t *testing.T) {
	input := `<html></html>
`

	actual, err := FromXML([]byte(input))
	assert.Nil(t, actual)
	assert.EqualError(t, err, "expected element type <urlset> but have <html>")
}

func TestFromXMLWithInvalidRandomString(t *testing.T) {
	input := `easy as 123`

	actual, err := FromXML([]byte(input))
	assert.Nil(t, actual)
	assert.EqualError(t, err, "EOF")
}
