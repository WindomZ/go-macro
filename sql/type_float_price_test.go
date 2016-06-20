package wgsqlx

import (
	"github.com/WindomZ/go-macro/json"
	"testing"
)

type testFloatPrice struct {
	Price1 FloatPrice `json:"price1"`
	Price2 FloatPrice `json:"price2"`
	Price3 FloatPrice `json:"price3"`
}

func TestJSONFloatPrice(t *testing.T) {
	SetFloatPricePrecision(5)
	p := &testFloatPrice{
		Price1: NewFloatPrice(1.012345),
		Price2: NewFloatPriceInt(201234),
		Price3: NewFloatPriceString("3.012345"),
	}
	t.Logf("%#v", p)
	data, err := gojson.Marshal(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", string(data))
	var pp testFloatPrice
	if err := gojson.Unmarshal(data, &pp); err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", pp)
}
