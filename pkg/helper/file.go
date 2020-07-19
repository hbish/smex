//Package helper helper
package helper

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
	"io/ioutil"

	"net/http"

	"github.com/pkg/errors"
)

// LoadSitemap - load sitemap currently using Remote to distinguish type
// TODO: enhance parsing logic to determine urls
func LoadSitemap(path string, Remote bool) ([]byte, error) {
	if Remote {
		return loadSitemapFromHTTP(path)
	}
	return loadSitemapFromFile(path)
}

// load sitemap given path
func loadSitemapFromFile(path string) ([]byte, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil || content == nil {
		return nil, errors.Wrap(err, "unable to read file content")
	}
	return content, nil
}

// load sitemap given url
func loadSitemapFromHTTP(url string) ([]byte, error) {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read http content")
	}
	return content, nil
}
