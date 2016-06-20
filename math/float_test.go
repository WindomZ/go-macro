package wdmath

import "testing"

func TestFloatPrecision(t *testing.T) {
	var f float64 = 2.012345
	t.Log(f)
	t.Log(FloatPrecision(f, 5, true))
}

func TestFloatRound(t *testing.T) {
	var f float64 = 2.012345
	t.Log(f)
	t.Log(FloatRound(f, 5))
}

func TestFloatFixed(t *testing.T) {
	var f float64 = 2.012345
	t.Log(f)
	t.Log(FloatFixed(f, 5))
}

func TestFloat(t *testing.T) {
	t.Log(1.012345, FloatPrecision(1.012345, 5, true), FloatRound(1.012345, 5), FloatFixed(1.012345, 5))
	t.Log(2.012345, FloatPrecision(2.012345, 5, true), FloatRound(2.012345, 5), FloatFixed(2.012345, 5))
	t.Log(3.012345, FloatPrecision(3.012345, 5, true), FloatRound(3.012345, 5), FloatFixed(3.012345, 5))
}
