package gojson

type JSON interface {
	JSONMarshal() interface{}
	JSONUnmarshal() interface{}
}
