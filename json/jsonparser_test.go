package gojson

import (
	"encoding/json"
	"github.com/buger/jsonparser"
	"github.com/fatih/structs"
	"testing"
)

type testLayer1 struct {
	testLayer2   `json:""`
	String1      string     `json:"string1"`
	Int1         int64      `json:"int1"`
	ArrayString1 []string   `json:"strings1"`
	Layer3       testLayer3 `json:""`
}

type testLayer2 struct {
	String2      string   `json:"string2"`
	Int2         int64    `json:"int2"`
	ArrayString2 []string `json:"strings2"`
}

type testLayer3 struct {
	String3      string   `json:"string3"`
	Int3         int64    `json:"int3"`
	ArrayString3 []string `json:"strings3"`
}

func TestJsonParser(t *testing.T) {
	s := &testLayer1{
		testLayer2: testLayer2{
			String2: "S2",
			Int2:    2,
			ArrayString2: []string{
				"Ss21",
				"Ss22",
				"Ss23",
			},
		},
		String1: "S1",
		Int1:    1,
		ArrayString1: []string{
			"Ss11",
			"Ss12",
			"Ss13",
		},
		Layer3: testLayer3{
			String3: "S3",
			Int3:    3,
			ArrayString3: []string{
				"Ss31",
				"Ss32",
				"Ss33",
			},
		},
	}
	if data, err := json.Marshal(s); err != nil {
		t.Fatal(err)
	} else {
		t.Log(string(data))
		//t.Logf("%#v", structs.Names(s))
		StructFields(s)
		jsonparser.EachKey(data, func(i int, value []byte, dataType jsonparser.ValueType, err error) {
			t.Logf("%#v", string(value))
		}, structs.Names(s))
	}
}
