package http

import (
	"fmt"
	"net/http"
)

func HEAD_KEY() string {
	return "Authorization"
}

func HEAD_VALUE(tokenString string) string {
	return fmt.Sprintf("Bearer %v", tokenString)
}

func SignedWithRequest(req *http.Request, tokenString string) *http.Request {
	req.Header.Set(HEAD_KEY(), HEAD_VALUE(tokenString))
	return req
}
