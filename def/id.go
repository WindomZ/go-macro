package def

import "github.com/WindomZ/go-macro/uuid"

func GetDefaultUUID() string {
	return uuid.NewSafeUUID()
}
