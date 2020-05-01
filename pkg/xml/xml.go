package xml

import "encoding/xml"

// Sitemap XML Protocol Implementation
// - info: https://www.sitemaps.org/protocol.html

// URLSet
type URLSet struct {
	XMLName xml.Name `xml:"urlset"`
	XMLNs   string   `xml:"xmlns,attr"`
	URL     []URL    `xml:"url"`
}

// URL is for every single location url
type URL struct {
	Loc        string  `xml:"loc"`
	LastMod    string  `xml:"lastmod,omitempty"`
	ChangeFreq string  `xml:"changefreq,omitempty"`
	Priority   float32 `xml:"priority,omitempty"`
}

func UnmarshalXML(rawXml []byte) (*URLSet, error) {
	urlSet := URLSet{}

	err := xml.Unmarshal(rawXml, &urlSet)

	if err != nil {
		return nil, err
	}

	return &urlSet, nil
}
