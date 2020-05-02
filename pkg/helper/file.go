package helper

import (
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

// load sitemap currently using Remote to distinguish type
// TODO: enhance parsing logic to determine urls
func LoadSitemap(path string, Remote bool) ([]byte, error) {
	if Remote {
		return loadSitemapFromHTTP(path)
	} else {
		return loadSitemapFromFile(path)
	}
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
	res, err := http.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "unable to retrieve sitemap")
	}
	defer res.Body.Close()
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read http content")
	}
	return content, nil
}
