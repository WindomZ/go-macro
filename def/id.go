package def

import (
	"fmt"
	"github.com/WindomZ/go-macro/uuid"
	"strings"
	"time"
)

func GetDefaultUUID() string {
	return uuid.NewSafeUUID()
}

func GetDefaultTimeID() string {
	return fmt.Sprintf("%v%02d", time.Now().Format("0601021504"), tv())
}

func GetDefaultOrderNo() string {
	return fmt.Sprintf("%v%02d", time.Now().Format("060102150405"), tv())
}

func GetDefaultOrderUUID() string {
	return strings.Replace(uuid.NewSafeUUID(), "-", "", -1)
}

func GetNumberId(id int64, n int) string {
	if n <= 0 {
		n = 64
	}
	s := fmt.Sprintf("%0"+fmt.Sprintf("%v", n)+"v", id)
	if len(s) > n {
		return s[len(s)-n:]
	}
	return s
}
