package gojson

import (
	"bytes"
	"encoding/json"
)

func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	return json.MarshalIndent(v, prefix, indent)
}

func MarshalTagHook(v interface{}, tag string, hook func(interface{}) interface{}) ([]byte, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	//fs := NewFields(v)
	dec := json.NewDecoder(bytes.NewReader(data))
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	var m map[string]interface{}
	if err := dec.Decode(&m); err != nil {
		return nil, err
	}
	if err := enc.Encode(&m); err != nil {
	}
	return data, nil
}
