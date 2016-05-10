package gomacro

import (
	"testing"
	"time"
)

func TestDateNextDayFirstTime(t *testing.T) {
	t.Logf("Now: %v", time.Now())
	t.Logf("NextDay: %v", DateNextDayFirstTime())
}

func TestDateNextDayLastTime(t *testing.T) {
	t.Logf("Now: %v", time.Now())
	t.Logf("NextDay: %v", DateNextDayLastTime())
}
