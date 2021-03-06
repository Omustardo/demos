// +build js

package assetloader

import (
	"io/ioutil"
	"net/http"
)

func loadFile(path string) ([]byte, error) {
	return httpGet(path)
}

// httpGet fetches the contents at the given url. This is used as a workaround for loading local assets while on the web.
// TODO: This prevents `gopherjs build` and then running locally because you can't use http.GET with local files. It has to be with http or https targets.
func httpGet(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
