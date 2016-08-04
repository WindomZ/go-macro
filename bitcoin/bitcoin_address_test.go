package bitcoin

import "testing"

func TestGenerateBitCoinAddress(t *testing.T) {
	m := make(map[string]bool)
	for i := 0; i < 100; i++ {
		address := GenerateBitCoinAddress()
		//t.Log(address)
		if _, ok := m[address]; ok {
			t.Fatalf("Duplicate address: %v", address)
		}
		m[address] = true
		if !VerifyBitCoinAddress(address) {
			t.Fatalf("Invalid address: %v", address)
		}
	}
}

func TestGenerateBitCoinAddressWith(t *testing.T) {
	m := make(map[string]bool)
	for i := 0; i < 100; i++ {
		address := GenerateBitCoinAddressWith([]byte("ABC"))
		//t.Log(address)
		if _, ok := m[address]; ok {
			t.Fatalf("Duplicate address: %v", address)
		}
		m[address] = true
		if !VerifyBitCoinAddress(address) {
			t.Fatalf("Invalid address: %v", address)
		}
	}
}

func TestGenerateBitCoinAddressWithPrefix(t *testing.T) {
	m := make(map[string]bool)
	for i := 0; i < 100; i++ {
		address := GenerateBitCoinAddressWithPrefix("ABCDF")
		//t.Log(address)
		if _, ok := m[address]; ok {
			t.Fatalf("Duplicate address: %v", address)
		}
	}
}
