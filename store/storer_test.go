package store

import (
	"strings"
	"testing"
)

func TestSaveFetch(t *testing.T) {
	id := "captcha id"
	d := string(randomDigits(10))
	Save(id, d)
	d2 := Fetch(id, false)
	if !strings.EqualFold(d, d2) {
		t.Errorf("saved %v, randomDigits returned got %v", d, d2)
	}
}
