package def

import (
	"fmt"
	"github.com/WindomZ/go-macro/uuid"
	"time"
)

func GetDefaultUUID() string {
	return uuid.NewSafeUUID()
}

func GetDefaultTimeID() string {
	return fmt.Sprintf("%v%02d", time.Now().Format("0601021504"), tv())
}
