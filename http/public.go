package http

import "net/http"

var defaultClient = NewHttpClient()

func HEAD_AUTH_KEY() string {
	return "Authorization"
}

func Post(url string, session string, data interface{}) (*http.Response, error) {
	defaultClient.Clear()
	if len(session) != 0 {
		defaultClient.SetHeader(HEAD_AUTH_KEY(), session)
	}
	return defaultClient.Post(url, data)
}

func Put(url string, session string, data interface{}) (*http.Response, error) {
	defaultClient.Clear()
	if len(session) != 0 {
		defaultClient.SetHeader(HEAD_AUTH_KEY(), session)
	}
	return defaultClient.Put(url, data)
}

func Get(url string, session string) (*http.Response, error) {
	defaultClient.Clear()
	if len(session) != 0 {
		defaultClient.SetHeader(HEAD_AUTH_KEY(), session)
	}
	return defaultClient.Get(url)
}

func Delete(url string, session string, data interface{}) (*http.Response, error) {
	defaultClient.Clear()
	if len(session) != 0 {
		defaultClient.SetHeader(HEAD_AUTH_KEY(), session)
	}
	return defaultClient.Delete(url, data)
}

func PostFile(url string, session string, filepath string) (*http.Response, error) {
	defaultClient.Clear()
	if len(session) != 0 {
		defaultClient.SetHeader(HEAD_AUTH_KEY(), session)
	}
	return defaultClient.PostFile(url, filepath)
}

func GetFile(url string, session string, filepath string) (*http.Response, error) {
	defaultClient.Clear()
	if len(session) != 0 {
		defaultClient.SetHeader(HEAD_AUTH_KEY(), session)
	}
	return defaultClient.GetFile(url, filepath)
}
