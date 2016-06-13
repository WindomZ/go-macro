package wgsqlx

import (
	"github.com/WindomZ/go-macro/json"
	"testing"
)

type testUnitPrice struct {
	Price1 UnitPrice `json:"price1"`
	Price2 UnitPrice `json:"price2"`
}

func TestJSONUnitPrice(t *testing.T) {
	SetUnitPricePrecision(5)
	p := &testUnitPrice{
		Price1: NewUnitPrice(101234),
		Price2: NewUnitPrice(201234),
	}
	t.Logf("%#v", p)
	data, err := gojson.Marshal(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", string(data))
	var pp testUnitPrice
	if err := gojson.Unmarshal(data, &pp); err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", pp)
}
