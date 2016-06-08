package wdmath

import "testing"

func TestFloatPrecision(t *testing.T) {
	var f float64 = 2.56789
	t.Log(f)
	t.Log(FloatPrecision(f, 2, true))
}
