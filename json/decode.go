package gojson

import (
	"encoding/json"
	"github.com/buger/jsonparser"
)

func Unmarshal(data []byte, v interface{}) error {
	if o, ok := v.(JSON); ok {
		return json.Unmarshal(data, o.JSONUnmarshal())
	}
	return json.Unmarshal(data, v)
}

func JSONGetString(data []byte, keys ...string) (string, error) {
	return jsonparser.GetString(data, keys...)
}

func JSONGetFloat(data []byte, keys ...string) (float64, error) {
	return jsonparser.GetFloat(data, keys...)
}

func JSONGetInt(data []byte, keys ...string) (int64, error) {
	return jsonparser.GetInt(data, keys...)
}

func JSONGetBoolean(data []byte, keys ...string) (bool, error) {
	return jsonparser.GetBoolean(data, keys...)
}
