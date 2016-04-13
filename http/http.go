package http

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strings"
)

var defaultClient = &http.Client{}

func Post(url string, session string, data interface{}) (*http.Response, error) {
	var body io.Reader = nil
	if data != nil {
		j, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		body = strings.NewReader(string(j))
	}
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if len(session) > 0 {
		req = SignedWithRequest(req, session)
	}
	return defaultClient.Do(req)
}

func Put(url string, session string, data interface{}) (*http.Response, error) {
	var body io.Reader = nil
	if data != nil {
		j, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		body = strings.NewReader(string(j))
	}
	req, err := http.NewRequest("PUT", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if len(session) > 0 {
		req = SignedWithRequest(req, session)
	}
	return defaultClient.Do(req)
}

func Get(url string, session string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	if len(session) > 0 {
		req = SignedWithRequest(req, session)
	}
	return defaultClient.Do(req)
}

func Delete(url string, session string, data interface{}) (*http.Response, error) {
	var body io.Reader = nil
	if data != nil {
		j, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		body = strings.NewReader(string(j))
	}
	req, err := http.NewRequest("DELETE", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if len(session) > 0 {
		req = SignedWithRequest(req, session)
	}
	return defaultClient.Do(req)
}

func PostFile(url string, session string, filepath string) (*http.Response, error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	fh, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer fh.Close()
	fileWriter, err := bodyWriter.CreateFormFile("file", path.Base(fh.Name()))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return nil, err
	}
	contentType := bodyWriter.FormDataContentType()
	err = bodyWriter.Close()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bodyBuf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	if len(session) > 0 {
		req = SignedWithRequest(req, session)
	}
	return defaultClient.Do(req)
}

func GetFile(url string, session string, filepath string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "multipart/form-data")
	if len(session) > 0 {
		req = SignedWithRequest(req, session)
	}
	resp, err := defaultClient.Do(req)
	defer resp.Body.Close()

	err = os.MkdirAll(path.Dir(filepath), 0766)
	if err != nil {
		return resp, err
	}
	f, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return resp, err
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	return resp, err
}
