package def

import (
	"fmt"
	"github.com/WindomZ/go-random/random"
	"time"
)

func GetDefaultUsername(tag string) string {
	if len(tag) == 0 {
		tag = "u"
	}
	return fmt.Sprintf("%v%v%v", tag, time.Now().Unix(), DiceValue())
}

func GetDefaultTel() string {
	return GetDefaultFormat("#tel")
}

func GetDefaultEmail() string {
	return GetDefaultFormat("#email")
}

func GetDefaultFormat(tag string) string {
	return fmt.Sprintf("%v-%v%v", tag, time.Now().Unix(), DiceValue())
}

func GetDefaultSalt(n int) string {
	return random.RandomString(n)
}
