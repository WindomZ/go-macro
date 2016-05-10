package gomacro

import (
	"testing"
	"time"
)

func TestDateTodayFirstTime(t *testing.T) {
	t.Logf("Now: %v", time.Now())
	t.Logf("Time: %v", DateTodayFirstTime())
}

func TestDateTodayLastTime(t *testing.T) {
	t.Logf("Now: %v", time.Now())
	t.Logf("Time: %v", DateTodayLastTime())
}

func TestDateTodayTime(t *testing.T) {
	t.Logf("Now: %v", time.Now())
	t.Logf("Time: %v", DateTodayTime(12, 22, 23, 0))
}

func TestDateNextDayFirstTime(t *testing.T) {
	t.Logf("Now: %v", time.Now())
	t.Logf("Time: %v", DateNextDayFirstTime())
}

func TestDateNextDayLastTime(t *testing.T) {
	t.Logf("Now: %v", time.Now())
	t.Logf("Time: %v", DateNextDayLastTime())
}
