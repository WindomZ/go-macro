package wgsqlx

import (
	"bytes"
	"database/sql/driver"
	"errors"
	"math"
	"strconv"
	"strings"
)

var (
	UnitPricePrecision int     = 2
	UnitPricePow       float64 = math.Pow10(UnitPricePrecision)
)

func SetUnitPricePrecision(e int) {
	UnitPricePrecision = e
	UnitPricePow = math.Pow10(UnitPricePrecision)
}

type UnitPrice int64

func NewUnitPrice(i int64) *UnitPrice {
	p := new(UnitPrice)
	p.SetInt64(i)
	return p
}

func (p *UnitPrice) MarshalJSON() ([]byte, error) {
	if p == nil {
		return nil, errors.New("MarshalJSON on nil pointer")
	}
	var b bytes.Buffer
	b.WriteByte('"')
	b.WriteString(p.StringFloat())
	b.WriteByte('"')
	return b.Bytes(), nil
}

func (p *UnitPrice) UnmarshalJSON(data []byte) error {
	if p == nil {
		return errors.New("UnmarshalJSON on nil pointer")
	} else if f, err := strconv.ParseFloat(strings.Replace(string(data), `"`, ``, -1), 64); err != nil {
		return err
	} else {
		p.SetFloat64(f)
	}
	return nil
}

func (p UnitPrice) Value() (driver.Value, error) {
	return p.Int64(), nil
}

func (p *UnitPrice) Scan(src interface{}) error {
	switch o := src.(type) {
	case int, int8, int16, int32, int64:
		*p = UnitPrice(o.(int64))
	case string:
		i, err := strconv.ParseInt(o, 10, 64)
		if err != nil {
			return err
		}
		*p = UnitPrice(i)
	case []byte:
		return p.Scan(string(o))
	default:
		return errors.New("Incompatible type for UnitPrice")
	}
	return nil
}

func (p UnitPrice) Int64() int64 {
	return int64(p)
}

func (p UnitPrice) SetInt64(i int64) {
	p = UnitPrice(i)
}

func (p UnitPrice) Float64() float64 {
	return float64(p) / UnitPricePow
}

func (p UnitPrice) SetFloat64(f float64) {
	p = UnitPrice(int64(f * UnitPricePow))
}

func (p UnitPrice) String() string {
	return strconv.FormatInt(int64(p), 10)
}

func (p UnitPrice) StringFloat() string {
	return strconv.FormatFloat(p.Float64(), 'f', UnitPricePrecision, 64)
}
