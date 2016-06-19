package wgsqlx

import (
	"github.com/WindomZ/go-macro/json"
	"testing"
)

type testIntPrice struct {
	Price1 IntPrice `json:"price1"`
	Price2 IntPrice `json:"price2"`
}

func TestJSONIntPrice(t *testing.T) {
	SetIntPricePrecision(5)
	p := &testIntPrice{
		Price1: NewIntPrice(101234),
		Price2: NewIntPrice(201234),
	}
	t.Logf("%#v", p)
	data, err := gojson.Marshal(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", string(data))
	var pp testIntPrice
	if err := gojson.Unmarshal(data, &pp); err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", pp)
}
