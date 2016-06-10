package http

import (
	"io/ioutil"
	"net/http"
)

func StringResponseBody(resp *http.Response) (string, error) {
	p, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	return string(p), err
}
