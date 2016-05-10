package gomacro

import "time"

func DateNextDayZeroTime() time.Time {
	y, m, d := time.Now().AddDate(0, 0, 1).Date()
	return time.Date(y, m, d, 0, 0, 0, 0, time.Local)
}
