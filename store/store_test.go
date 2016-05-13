package store

import (
	"bytes"
	"fmt"
	"testing"
	"time"
)

const (
	testCollectNum = 100
	testExpiration = 10 * time.Minute
)

var testIdx byte = 0

func randomId() string {
	testIdx++
	return fmt.Sprintf("%v%v", time.Now().Unix(), testIdx)
}

func randomDigits(length int) []byte {
	testIdx++
	b := []byte(fmt.Sprintf("%v%v", time.Now().Unix(), testIdx))
	r := make([]byte, length)
	if length > len(b) {
		length = len(b)
	}
	for i := 0; i < length; i++ {
		r[i] = b[i]
	}
	return r
}

func TestSetGet(t *testing.T) {
	s := NewMemoryStore(testCollectNum, testExpiration)
	id := "captcha id"
	d := randomDigits(10)
	s.Set(id, d)
	d2 := s.Get(id, false)
	if d2 == nil || !bytes.Equal(d, d2) {
		t.Errorf("saved %v, randomDigits returned got %v", d, d2)
	}
}

func TestGetClear(t *testing.T) {
	s := NewMemoryStore(testCollectNum, testExpiration)
	id := "captcha id"
	d := randomDigits(10)
	s.Set(id, d)
	d2 := s.Get(id, true)
	if d2 == nil || !bytes.Equal(d, d2) {
		t.Errorf("saved %v, getDigitsClear returned got %v", d, d2)
	}
	d2 = s.Get(id, false)
	if d2 != nil {
		t.Errorf("getDigitClear didn't clear (%q=%v)", id, d2)
	}
}

func TestCollect(t *testing.T) {
	s := NewMemoryStore(10, -1)
	ids := make([]string, 10)
	d := randomDigits(10)
	for i := range ids {
		ids[i] = randomId()
		s.Set(ids[i], d)
	}
	s.(*memoryStore).collect()
	nc := 0
	for i := range ids {
		d2 := s.Get(ids[i], false)
		if d2 != nil {
			t.Errorf("%d: not collected", i)
			nc++
		}
	}
	if nc > 0 {
		t.Errorf("= not collected %d out of %d captchas", nc, len(ids))
	}
}

func BenchmarkSetCollect(b *testing.B) {
	b.StopTimer()
	d := randomDigits(10)
	s := NewMemoryStore(9999, -1)
	ids := make([]string, 1000)
	for i := range ids {
		ids[i] = randomId()
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 1000; j++ {
			s.Set(ids[j], d)
		}
		s.(*memoryStore).collect()
	}
}
