package def

import "testing"

func TestGetDefaultUsername(t *testing.T) {
	t.Log(GetDefaultUsername(""))
	t.Log(GetDefaultUsername("test"))
}

func TestGetDefaultTel(t *testing.T) {
	t.Log(GetDefaultTel())
}

func TestGetDefaultEmail(t *testing.T) {
	t.Log(GetDefaultEmail())
}
