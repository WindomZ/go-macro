package gomacro

import "time"

func DateFirstTime(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.Local)
}

func DateLastTime(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 23, 59, 59, 999999999, time.Local)
}

func DateTodayFirstTime() time.Time {
	return DateFirstTime(time.Now().Date())
}

func DateTodayLastTime() time.Time {
	return DateLastTime(time.Now().Date())
}

func DateTodayTime(hour, min, sec, nsec int) time.Time {
	year, month, day := time.Now().Date()
	return time.Date(year, month, day, hour, min, sec, nsec, time.Local)
}

func DateNextDayFirstTime() time.Time {
	return DateFirstTime(time.Now().AddDate(0, 0, 1).Date())
}

func DateNextDayLastTime() time.Time {
	return DateLastTime(time.Now().AddDate(0, 0, 1).Date())
}
