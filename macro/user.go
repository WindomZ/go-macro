package macro

const MIN_LEN_USERNAME int = 4
const MAX_LEN_USERNAME int = 64

func ValidUsername(name string) bool {
	l := len(name)
	return (l >= MIN_LEN_USERNAME && l <= MAX_LEN_USERNAME)
}

const MIN_LEN_PASSWORD int = 6
const MAX_LEN_PASSWORD int = 32

func ValidPassword(pwd string) bool {
	l := len(pwd)
	return (l >= MIN_LEN_PASSWORD && l <= MAX_LEN_PASSWORD)
}

func ValidBitCoinUsername(name string) bool {
	return len(name) == 34
}
