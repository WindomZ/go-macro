package gojson

import "errors"

var (
	ErrInterface error = errors.New("gojson: unsupported interface")
	ErrType            = errors.New("gojson: unsupported type")
)
