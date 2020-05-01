package helper

import (
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

func LoadSitemap(path string, isRemote bool) ([]byte, error) {
	if isRemote {
		return loadSitemapFromHTTP(path)
	} else {
		return loadSitemapFromFile(path)
	}
}

func loadSitemapFromFile(path string) ([]byte, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil || content == nil {
		return nil, errors.Wrap(err, "unable to read file content")
	}
	return content, nil
}

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
