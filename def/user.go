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
	return fmt.Sprintf("%v%v%v", tag, time.Now().Unix(), tv10())
}

func GetDefaultTel() string {
	return GetDefaultFormat("#tel")
}

func GetDefaultEmail() string {
	return GetDefaultFormat("#email")
}

func GetDefaultCert() string {
	return GetDefaultFormat("#cert")
}

func GetDefaultBank() string {
	return GetDefaultFormat("#bank")
}

func GetDefaultFormat(tag string) string {
	return fmt.Sprintf("%v-%v%v", tag, time.Now().Unix(), tv10())
}

func GetDefaultSalt(n int) string {
	return random.RandomString(n)
}
