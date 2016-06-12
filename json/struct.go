package gojson

import (
	"fmt"
	"github.com/fatih/structs"
	"reflect"
)

type Fields struct {
	Object     interface{}
	FieldsName []string
	FieldsMap  map[string]*structs.Field
}

func NewFields(v interface{}) *Fields {
	f := &Fields{
		Object:     v,
		FieldsName: make([]string, 0, 10),
		FieldsMap:  make(map[string]*structs.Field),
	}
	return f.fields(structs.Fields(v))
}

func StructFields(v interface{}) {
	fs := NewFields(v).FieldsMap
	for _, f := range fs {
		println("Result: " + f.Name())
	}
}

func (f *Fields) fields(fs []*structs.Field, tags ...string) *Fields {
	var tag, key string
	if len(tags) != 0 {
		tag = tags[0]
	}
	for _, ff := range fs {
		if len(tag) == 0 {
			key = ff.Name()
		} else {
			key = fmt.Sprintf("%v::%v", tag, ff.Name())
		}
		//println("ff: ", ff.Name(), ff.IsEmbedded(), ff.IsExported(), (ff.Kind() == reflect.Struct))
		if (ff.IsEmbedded() || ff.IsExported()) && ff.Kind() == reflect.Struct {
			if ffs := ff.Fields(); len(ffs) != 0 {
				f.fields(ffs, key)
				continue
			}
		}
		f.FieldsName = append(f.FieldsName, key)
		f.FieldsMap[key] = ff
	}
	return f
}
