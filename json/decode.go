package gojson

import "encoding/json"

func Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
