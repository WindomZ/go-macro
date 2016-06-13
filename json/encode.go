package gojson

import "encoding/json"

func Marshal(v interface{}) ([]byte, error) {
	if o, ok := v.(JSON); ok {
		return json.Marshal(o.JSONMarshal())
	}
	return json.Marshal(v)
}

func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	if o, ok := v.(JSON); ok {
		return json.MarshalIndent(o.JSONMarshal(), prefix, indent)
	}
	return json.MarshalIndent(v, prefix, indent)
}
