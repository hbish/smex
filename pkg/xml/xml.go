// Package xml xml
package xml

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
	"bytes"
	"encoding/xml"
	"regexp"
)

// Sitemap XML Protocol Implementation
// - info: https://www.sitemaps.org/protocol.html

// URLSet urlSet struct
type URLSet struct {
	XMLName xml.Name `xml:"urlset"`
	XMLNs   string   `xml:"xmlns,attr"`
	URL     []URL    `xml:"url"`
}

// URL url struct
type URL struct {
	Loc        string  `xml:"loc" json:"loc"`
	LastMod    string  `xml:"lastmod,omitempty" json:"lastmod,omitempty"`
	ChangeFreq string  `xml:"changefreq,omitempty" json:"changefreq,omitempty"`
	Priority   float32 `xml:"priority,omitempty" json:"priority,omitempty"`
	Image      []Image `xml:"image,omitempty" json:"image,omitempty"`
	Video      []Video `xml:"video,omitempty" json:"video,omitempty"`
	News       []News  `xml:"news,omitempty" json:"news,omitempty"`
}

// Image image struct
type Image struct {
	Loc         string `xml:"loc" json:"loc"`
	Title       string `xml:"title,omitempty" json:"title,omitempty"`
	Caption     string `xml:"caption,omitempty" json:"caption,omitempty"`
	GeoLocation string `xml:"geo_location,omitempty" json:"geo_location,omitempty"`
	License     string `xml:"license,omitempty" json:"license,omitempty"`
}

// Video video struct
type Video struct {
	ThumbnailLoc         string   `xml:"thumbnail_location" json:"thumbnail_location"`
	Title                string   `xml:"title,omitempty" json:"title,omitempty"`
	Description          string   `xml:"description,omitempty" json:"description,omitempty"`
	ContentLoc           string   `xml:"content_loc,omitempty" json:"content_loc,omitempty"`
	PlayerLoc            string   `xml:"player_loc,omitempty" json:"player_loc,omitempty"`
	Duration             int      `xml:"duration,omitempty" json:"duration,omitempty"`
	ExpirationDate       string   `xml:"expiration_date,omitempty" json:"expiration_date,omitempty"`
	Rating               float32  `xml:"rating,omitempty" json:"rating,omitempty"`
	ViewCount            int64    `xml:"view_count,omitempty" json:"view_count,omitempty"`
	PublicationDate      string   `xml:"publication_date,omitempty" json:"publication_date,omitempty"`
	FamilyFriendly       string   `xml:"family_friendly,omitempty" json:"family_friendly,omitempty"`
	Restriction          string   `xml:"restriction,omitempty" json:"restriction,omitempty"`
	Platform             string   `xml:"platform,omitempty" json:"platform,omitempty"`
	Price                float32  `xml:"price,omitempty" json:"price,omitempty"`
	RequiresSubscription string   `xml:"requires_subscription,omitempty" json:"requires_subscription,omitempty"`
	Uploader             string   `xml:"uploader,omitempty" json:"uploader,omitempty"`
	Live                 string   `xml:"live,omitempty" json:"live,omitempty"`
	Tag                  []string `xml:"tag,omitempty" json:"tag,omitempty"`
	Category             string   `xml:"category,omitempty" json:"category,omitempty"`
}

// News news struct
type News struct {
	Publication     Publication `xml:"publication,omitempty" json:"publication,omitempty"`
	PublicationDate string      `xml:"publication_date,omitempty" json:"publication_date,omitempty"`
	Title           string      `xml:"title,omitempty" json:"title,omitempty"`
	Genres          string      `xml:"genres,omitempty" json:"genres,omitempty"`
	Keywords        string      `xml:"keywords,omitempty" json:"keywords,omitempty"`
	Access          string      `xml:"access,omitempty" json:"access,omitempty"`
	StockTickers    string      `xml:"stock_tickers,omitempty" json:"stock_tickers,omitempty"`
}

// Publication publication struct
type Publication struct {
	Name     string `xml:"name,omitempty" json:"name,omitempty"`
	Language string `xml:"language,omitempty" json:"language,omitempty"`
}

// unmarshalXML unmarshal raw data
func unmarshalXML(rawXML []byte) (*URLSet, error) {
	urlSet := URLSet{}

	// validate xml without storing
	if err := xml.Unmarshal(rawXML, new(interface{})); err != nil {
		return nil, err
	}

	// decode xml and trim white spaces
	reader := bytes.NewReader(rawXML)
	d := xml.NewDecoder(reader)
	td := xml.NewTokenDecoder(TrimmingTokenReader{d})
	err := td.Decode(&urlSet)

	if err != nil {
		return nil, err
	}

	return &urlSet, nil
}

// UnmarshalXMLP unmarshal xml and filter by pattern
func UnmarshalXMLP(rawXML []byte, pattern string) (*URLSet, error) {
	urlSet, err := unmarshalXML(rawXML)
	if err != nil {
		return nil, err
	}

	if pattern != "" {
		regex, err := regexp.Compile(pattern)
		if err != nil {
			return nil, err
		}

		var filteredUrls []URL
		for _, url := range urlSet.URL {
			matched := regex.MatchString(url.Loc)
			if matched {
				filteredUrls = append(filteredUrls, url)
			}
		}

		urlSet.URL = filteredUrls
	}

	return urlSet, nil
}

// TrimmingTokenReader Trimming TokenReader
type TrimmingTokenReader struct {
	dec *xml.Decoder
}

// Token Trimming token
func (tr TrimmingTokenReader) Token() (xml.Token, error) {
	t, err := tr.dec.Token()
	if cd, ok := t.(xml.CharData); ok {
		t = xml.CharData(bytes.TrimSpace(cd))
	}
	return t, err
}
