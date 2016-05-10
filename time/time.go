package gomacro

import "time"

func DateNextDayFirstTime() time.Time {
	y, m, d := time.Now().AddDate(0, 0, 1).Date()
	return time.Date(y, m, d, 0, 0, 0, 0, time.Local)
}

func DateNextDayLastTime() time.Time {
	y, m, d := time.Now().AddDate(0, 0, 1).Date()
	return time.Date(y, m, d, 23, 59, 59, 999999999, time.Local)
}
