package store

import "time"

var _store Store

func InitStore(collectNum int, expiration time.Duration) {
	_store = NewMemoryStore(collectNum, expiration)
}

func GetStore() Store {
	if _store == nil {
		InitStore(100, 10*time.Minute)
	}
	return _store
}

func Save(id string, str string) {
	GetStore().Set(id, []byte(str))
}

func Fetch(id string, clear bool) string {
	b := GetStore().Get(id, clear)
	if b == nil {
		return ""
	}
	return string(b)
}
