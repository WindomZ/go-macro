package gomacro

import (
	"testing"
	"time"
)

func TestDateFirstTime(t *testing.T) {
	t.Logf("Now: %v", time.Now())
	t.Logf("Time: %v", DateFirstTime())
}

func TestDateLastTime(t *testing.T) {
	t.Logf("Now: %v", time.Now())
	t.Logf("Time: %v", DateLastTime())
}

func TestDateNextDayFirstTime(t *testing.T) {
	t.Logf("Now: %v", time.Now())
	t.Logf("Time: %v", DateNextDayFirstTime())
}

func TestDateNextDayLastTime(t *testing.T) {
	t.Logf("Now: %v", time.Now())
	t.Logf("Time: %v", DateNextDayLastTime())
}
