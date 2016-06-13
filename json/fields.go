package gojson

import (
	"fmt"
	"github.com/fatih/structs"
	"reflect"
	"strings"
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

func (f *Fields) Names() [][]string {
	rs := make([][]string, 0, len(f.FieldsName))
	for _, n := range f.FieldsName {
		rs = append(rs, strings.Split(n, "::"))
	}
	return rs
}

func (f *Fields) FieldIndex(idx int) (*structs.Field, bool) {
	if idx < 0 || len(f.FieldsName) <= idx {
	} else if ff, ok := f.FieldsMap[f.FieldsName[idx]]; ok {
		return ff, true
	}
	return nil, false
}

func (f *Fields) Field(name ...string) (*structs.Field, bool) {
	key := strings.Join(name, "::")
	if ff, ok := f.FieldsMap[key]; ok {
		return ff, true
	}
	return nil, false
}

func (f *Fields) FieldsTag(tag string) ([]*structs.Field, []string, bool) {
	fs := make([]*structs.Field, 0, len(f.FieldsMap))
	vs := make([]string, 0, len(f.FieldsMap))
	for _, ff := range f.FieldsMap {
		if v := ff.Tag(tag); len(v) != 0 {
			fs = append(fs, ff)
			vs = append(vs, v)
		}
	}
	return fs, vs, len(fs) != 0
}

func (f *Fields) FieldTagValue(tag, value string) (*structs.Field, bool) {
	for _, ff := range f.FieldsMap {
		if v := ff.Tag(tag); v == value {
			return ff, true
		}
	}
	return nil, false
}
