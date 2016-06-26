package def

import "testing"

func TestGetDefaultUUID(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(GetDefaultUUID())
	}
}

func TestGetNumberId(t *testing.T) {
	t.Log(GetNumberId(1500, 10))
	t.Log(GetNumberId(1500, 3))
}

func TestGetDefaultOrderUUID(t *testing.T) {
	t.Log(GetDefaultOrderUUID())
}

func TestGetDefaultTimeID(t *testing.T) {
	for i := 0; i < 100; i++ {
		t.Log(GetDefaultTimeID())
	}
}

func TestGetDefaultTimeId(t *testing.T) {
	for i := 0; i < 100; i++ {
		t.Log(GetDefaultTimeId())
	}
}

func TestGetDefaultOrderNo(t *testing.T) {
	for i := 0; i < 100; i++ {
		t.Log(GetDefaultOrderNo())
	}
}
