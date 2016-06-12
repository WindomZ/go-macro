package gojson

import (
	"encoding/json"
	"errors"
	"github.com/fatih/structs"
)

var (
	ErrInterface error = errors.New("gojson: unsupported interface")
	ErrType            = errors.New("gojson: unsupported type")
)

func Marshal(v interface{}) ([]byte, error) {
	if !structs.IsStruct(v) {
		return nil, ErrInterface
	}
	fs := structs.Fields(v)
	if len(fs) == 0 {
		return json.Marshal(v)
	}
	for _, f := range fs {
		f.IsEmbedded()
	}
	return nil, nil
}

func Unmarshal(data []byte, v interface{}) error {
	return nil
}
