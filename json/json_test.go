package gojson

import (
	"bytes"
	"encoding/json"
	"testing"
)

type testLayer1 struct {
	testLayer2   `json:""`
	String1      string     `json:"string1"`
	Int1         int64      `json:"int1,string"`
	ArrayString1 []string   `json:"strings1"`
	Layer3       testLayer3 `json:""`
}

type testLayer2 struct {
	String2      string   `json:"string2"`
	Int2         int64    `json:"int2,string"`
	ArrayString2 []string `json:"strings2"`
}

type testLayer3 struct {
	String3      string   `json:"string3"`
	Int3         int64    `json:"int3,string"`
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
		////t.Logf("%#v", structs.Names(s))
		//jsonparser.EachKey(data, func(i int, value []byte, dataType jsonparser.ValueType, err error) {
		//	t.Logf("%#v", string(value))
		//}, structs.Names(s))
		var ss testLayer1
		if err := json.Unmarshal(data, &ss); err != nil {
			t.Fatal(err)
		}
		t.Logf("%#v", ss)

		dec := json.NewDecoder(bytes.NewReader(data))
		var buf bytes.Buffer
		enc := json.NewEncoder(&buf)
		var v map[string]interface{}
		if err := dec.Decode(&v); err != nil {
			t.Fatal(err)
		}
		t.Logf("%#v", v)
		if err := enc.Encode(&v); err != nil {
			t.Fatal(err)
		}
		t.Logf("%#v", string(buf.Bytes()))
	}

}
