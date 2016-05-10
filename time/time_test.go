package gomacro

import (
	"testing"
	"time"
)

func TestDateNextDayZeroTime(t *testing.T) {
	t.Logf("Now: %v", time.Now())
	t.Logf("NextDay: %v", DateNextDayZeroTime())
}
