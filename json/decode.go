package gojson

import "encoding/json"

func Unmarshal(data []byte, v interface{}) error {
	if o, ok := v.(JSON); ok {
		return json.Unmarshal(data, o.JSONUnmarshal())
	}
	return json.Unmarshal(data, v)
}
