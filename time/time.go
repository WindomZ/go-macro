package gomacro

import "time"

func firstTime(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.Local)
}

func lastTime(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 23, 59, 59, 999999999, time.Local)
}

func DateFirstTime() time.Time {
	return firstTime(time.Now().Date())
}

func DateLastTime() time.Time {
	return lastTime(time.Now().Date())
}

func DateNextDayFirstTime() time.Time {
	return firstTime(time.Now().AddDate(0, 0, 1).Date())
}

func DateNextDayLastTime() time.Time {
	return lastTime(time.Now().AddDate(0, 0, 1).Date())
}
