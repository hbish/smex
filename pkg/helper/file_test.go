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
package helper

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

var fs = afero.NewMemMapFs()
var (
	mux    *http.ServeMux
	server *httptest.Server
)

func TestLoadSitemap_Local(t *testing.T) {
	// given
	_, _ = fs.Create(t.Name())
	_ = afero.WriteFile(fs, t.Name(), []byte("testing"), 0644)

	// when
	actual, _ := LoadSitemap(t.Name(), false)

	// then
	expected, _ := ioutil.ReadFile(t.Name())
	assert.Equal(t, string(expected), string(actual))
}

func TestLoadSitemap_Local_NotFound(t *testing.T) {
	// given
	// when
	_, err := LoadSitemap(t.Name(), false)

	// then
	assert.EqualError(t, err, "unable to read file content: open TestLoadSitemap_Local_NotFound: no such file or directory")
}

func TestLoadSitemap_Remote(t *testing.T) {
	// given
	xml := "<?xml version=\"1.0\" encoding=\"UTF-8\"?>" +
		"<urlset xmlns=\"http://www.sitemaps.org/schemas/sitemap/0.9\">" +
		"<url><loc>http://www.example.com/</loc><lastmod>2005-01-01</lastmod><changefreq>monthly</changefreq><priority>0.8</priority></url>" +
		"</urlset>"
	mux = http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(xml))
	})
	server = httptest.NewServer(mux)

	// when
	actual, _ := LoadSitemap(server.URL+"/sitemap.xml", true)

	// then
	assert.Equal(t, xml, string(actual))
}

func TestLoadSitemap_Remote_NotFound(t *testing.T) {
	// given
	mux = http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})
	server = httptest.NewServer(mux)

	// when
	actual, err := LoadSitemap(server.URL+"/sitemap.xml", true)

	// then
	assert.NoError(t, err)
	assert.Equal(t, "", string(actual))
}

func TestLoadSitemap_Remote_GatewayError(t *testing.T) {
	// given
	// when
	_, err := LoadSitemap("http://127.0.0.1:8999/sitemap.xml", true)

	// then
	assert.EqualError(t, err, "Get \"http://127.0.0.1:8999/sitemap.xml\": dial tcp 127.0.0.1:8999: connect: connection refused")
}
