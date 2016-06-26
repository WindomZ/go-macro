package macro

import "strings"

const EMPTY_ID string = "00000000-0000-0000-0000-000000000000"

func IsEmpty(id string) bool {
	return strings.EqualFold(id, EMPTY_ID)
}

func ToId(id string) string {
	return strings.ToLower(strings.TrimSpace(id))
}

func NoDashId(id string) string {
	return strings.Replace(ToId(id), "-", "", -1)
}

func ValidId(id string) bool {
	if IsEmpty(id) {
		return false
	} else if len(id) == 32 {
		return true
	} else if len(id) == 36 && strings.Count(id, "-") == 4 {
		return true
	}
	return false
}

func ValidTimeID(id string) bool {
	return len(id) == 12
}

func ValidTimeId(id string) bool {
	return len(id) == 16
}

func ValidOrderNo(id string) bool {
	return len(id) == 14
}
