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
	IntPricePrecision int     = 2
	IntPricePow       float64 = math.Pow10(IntPricePrecision)
)

func SetIntPricePrecision(e int) {
	IntPricePrecision = e
	IntPricePow = math.Pow10(IntPricePrecision)
}

type IntPrice int64

func NewIntPrice(i int64) IntPrice {
	return IntPrice(i)
}

func NewIntPriceFloat(i float64) IntPrice {
	return IntPrice(int64(i * IntPricePow))
}

func (p *IntPrice) MarshalJSON() ([]byte, error) {
	if p == nil {
		return nil, errors.New("MarshalJSON on nil pointer")
	}
	var b bytes.Buffer
	b.WriteByte('"')
	b.WriteString(p.StringFloat())
	b.WriteByte('"')
	return b.Bytes(), nil
}

func (p *IntPrice) UnmarshalJSON(data []byte) error {
	if p == nil {
		return errors.New("UnmarshalJSON on nil pointer")
	} else if f, err := strconv.ParseFloat(strings.Replace(string(data), `"`, ``, -1), 64); err != nil {
		return err
	} else {
		p.SetFloat64(f)
	}
	return nil
}

func (p IntPrice) Value() (driver.Value, error) {
	return p.Int64(), nil
}

func (p *IntPrice) Scan(src interface{}) error {
	switch o := src.(type) {
	case int, int8, int16, int32, int64:
		*p = IntPrice(o.(int64))
	case string:
		i, err := strconv.ParseInt(o, 10, 64)
		if err != nil {
			return err
		}
		*p = IntPrice(i)
	case []byte:
		return p.Scan(string(o))
	default:
		return errors.New("Incompatible type for IntPrice")
	}
	return nil
}

func (p IntPrice) Int64() int64 {
	return int64(p)
}

func (p *IntPrice) SetInt64(i int64) *IntPrice {
	*p = IntPrice(i)
	return p
}

func (p IntPrice) Float64() float64 {
	return float64(p) / IntPricePow
}

func (p *IntPrice) SetFloat64(f float64) *IntPrice {
	*p = IntPrice(int64(f * IntPricePow))
	return p
}

func (p IntPrice) String() string {
	return strconv.FormatInt(int64(p), 10)
}

func (p IntPrice) StringFloat() string {
	return strconv.FormatFloat(p.Float64(), 'f', IntPricePrecision, 64)
}

func (p *IntPrice) IsPositive() bool {
	return p.Int64() > 0
}

func (p *IntPrice) IsNegative() bool {
	return p.Int64() < 0
}

func (p *IntPrice) Zero() {
	p.SetInt64(0)
}

func (p *IntPrice) IsZero() bool {
	return p.Int64() == 0
}

func (p *IntPrice) Equal(f float64) bool {
	return p.Float64() == f
}

func (p *IntPrice) LessThan(f float64) bool {
	return p.Float64() < f
}

func (p *IntPrice) LessEqual(f float64) bool {
	return p.Float64() <= f
}

func (p *IntPrice) GreaterThan(f float64) bool {
	return p.Float64() > f
}

func (p *IntPrice) GreaterEqual(f float64) bool {
	return p.Float64() >= f
}

// rounded p+q and returns p
func (p *IntPrice) Add(q IntPrice) *IntPrice {
	return p.SetInt64(p.Int64() + q.Int64())
}

// rounded p-q and returns p
func (p *IntPrice) Sub(q IntPrice) *IntPrice {
	return p.SetInt64(p.Int64() - q.Int64())
}

// rounded product p*q and returns p
func (p *IntPrice) Mul(q IntPrice) *IntPrice {
	return p.SetInt64(p.Int64() * q.Int64())
}

// rounded quotient p/q and returns p
func (p *IntPrice) Quo(q IntPrice) *IntPrice {
	return p.SetInt64(p.Int64() / q.Int64())
}

// rounded p+x... and returns p
func (p *IntPrice) Sum(x ...IntPrice) *IntPrice {
	for _, y := range x {
		p.Add(y)
	}
	return p
}

// rounded p-x... and returns p
func (p *IntPrice) Diff(x ...IntPrice) *IntPrice {
	for _, y := range x {
		p.Sub(y)
	}
	return p
}

// rounded p+x... and returns
func (p IntPrice) GetSum(x ...IntPrice) IntPrice {
	sum := p.Int64()
	for _, y := range x {
		sum += y.Int64()
	}
	return NewIntPrice(sum)
}

// rounded p+x... and returns
func (p IntPrice) GetDiff(x ...IntPrice) IntPrice {
	diff := p.Int64()
	for _, y := range x {
		diff -= y.Int64()
	}
	return NewIntPrice(diff)
}

// returns negation
func (p IntPrice) GetNegation() IntPrice {
	return NewIntPrice(-p.Int64())
}

// Cmp compares p and p and returns:
//
//   -1 if p <  p
//    0 if p == p (incl. -0 == 0, -Inf == -Inf, and +Inf == +Inf)
//   +1 if p >  p
//
func (p IntPrice) Cmp(q IntPrice) int {
	switch {
	case p.Int64() < q.Int64():
		return -1
	case p.Int64() > q.Int64():
		return +1
	}
	return 0
}
