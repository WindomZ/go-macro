package def

import "testing"

func TestGetDefaultUUID(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(GetDefaultUUID())
	}
}
