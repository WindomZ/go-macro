package store

import (
	"strings"
	"testing"
)

func TestSaveFetch(t *testing.T) {
	id := "captcha id"
	d := string(randomDigits(10))
	Save(id, d)
	if f := Fetch(id, false); !strings.EqualFold(d, f) {
		t.Errorf("saved %v, randomDigits returned got %v", d, f)
	}
	if f := Fetch(id, false); !strings.EqualFold(d, f) {
		t.Errorf("saved %v, randomDigits returned got %v", d, f)
	}
	Remove(id)
	if f := Fetch(id, false); strings.EqualFold(d, f) {
		t.Errorf("saved %v, randomDigits returned got %v", d, f)
	}
}
