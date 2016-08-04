package bitcoin

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/WindomZ/go-random/random"
	"math/big"
	"sync"
)

var (
	idx   byte
	mutex *sync.Mutex
)

func init() {
	idx = 0
	mutex = &sync.Mutex{}
}

func GenerateBitCoinAddress() string {
	return newSafeBitCoinAddress(nil)
}

func GenerateBitCoinAddressWith(data []byte) string {
	return newSafeBitCoinAddress(data)
}

func GenerateBitCoinAddressWithPrefix(pre string) string {
	add := GenerateBitCoinAddressWith([]byte(pre))
	return pre + add[len(pre):]
}

func newSafeBitCoinAddress(data []byte) string {
	mutex.Lock()
	defer mutex.Unlock()
	idx++
	return makeBitCoinAddress("00", data, idx)
}

func VerifyBitCoinAddress(s string) bool {
	return VerifyBitCoinAddressWith(s, nil)
}

func VerifyBitCoinAddressWith(s string, data []byte) bool {
	if b, err := checkDecodeString(s); err != nil {
		return false
	} else if data != nil && len(data) != 0 {
		if len(b)-kMinPrefixLength < len(data) {
			return false
		} else {
			for i := 0; i < len(data); i++ {
				if b[i+kMinPrefixLength] != data[i] {
					return false
				}
			}
		}
	}
	return true
}

var base58alphabet = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

var base58reverse []int

func init() {
	base58reverse = make([]int, 256)
	for i, _ := range base58reverse {
		base58reverse[i] = -1
	}
	for i, c := range base58alphabet {
		base58reverse[int(c)] = i
	}
}

func reverseInplace(data []byte) {
	for i := 0; i < len(data)/2; i++ {
		tmp := data[i]
		data[i] = data[len(data)-i-1]
		data[len(data)-i-1] = tmp
	}
}

func encode(dst, src []byte) int {
	zeros := 0
	for _, b := range src {
		if int(b) == 0 {
			zeros++
		} else {
			break
		}
	}
	i := new(big.Int).SetBytes(src)
	big58 := big.NewInt(58)
	big0 := big.NewInt(0)

	var index int
	for i.Cmp(big0) > 0 {
		tmp := new(big.Int).Mod(i, big58)
		i.Div(i, big58)
		dst[index] = base58alphabet[tmp.Int64()]
		index++
	}
	for ; zeros > 0; zeros-- {
		dst[index] = base58alphabet[0]
		index++
	}
	reverseInplace(dst[0:index])
	return index
}

func checkEncode(dst, src []byte) int {
	var buf bytes.Buffer
	buf.Write(src)
	buf.Write(checksum(buf.Bytes()))
	return encode(dst, buf.Bytes())
}

func encodedMaxLen(x int) int {
	return 2 * x
}

const (
	kMinPrefixLength = 1
	kMaxPrefixLength = 4
	kChecksumLength  = 4
)

func checkEncodedMaxLen(x int) int {
	return encodedMaxLen(x + kChecksumLength + kMaxPrefixLength)
}

func checkEncodeToString(src []byte) string {
	dst := make([]byte, checkEncodedMaxLen(len(src)))
	n := checkEncode(dst, src)
	return string(dst[0:n])
}

func checksum(data []byte) []byte {
	first := sha256.Sum256(data)
	second := sha256.Sum256(first[:])
	return second[0:kChecksumLength]
}

func makeBitCoinAddress(head string, pre []byte, idx byte) string {
	bb := bytes.Buffer{}
	start := 0
	if pre != nil && len(pre) > 0 && len(pre) < 12 {
		start = len(pre)
		bb.Write(pre)
	}
	bb.Write(random.RandomTimeBytes(20 - start))
	data := bb.Bytes()
	data[byte(start+8)+idx%(12-byte(start))] = idx
	h, _ := hex.DecodeString(head)
	return checkEncodeToString(append(h, data...))
}

func decodedMaxLen(x int) int {
	return x
}

func decode(dst, src []byte) (int, error) {
	zeros := 0
	for _, b := range src {
		if b == base58alphabet[0] {
			zeros++
		} else {
			break
		}
	}
	big58 := big.NewInt(58)
	i := big.NewInt(0)
	for _, c := range src {
		i.Mul(i, big58)
		val := base58reverse[int(c)]
		if val < 0 {
			return 0, fmt.Errorf("Bad character %q in input.", c)
		}
		i.Add(i, big.NewInt(int64(val)))
	}
	length := len(i.Bytes()) + zeros
	if len(dst) < length {
		return 0, fmt.Errorf("Destination buffer not big enough; "+
			"need %d, got %d.", length, len(dst))
	}
	copy(dst[zeros:length], i.Bytes())
	return length, nil
}

func verifyChecksum(data []byte) bool {
	check := data[len(data)-kChecksumLength : len(data)]
	if !bytes.Equal(check, checksum(data[0:len(data)-kChecksumLength])) {
		return false
	}
	return true
}

func checkDecodeString(s string) ([]byte, error) {
	if len(s) < kChecksumLength+kMinPrefixLength {
		return nil, fmt.Errorf("Input too short.")
	}
	dst := make([]byte, decodedMaxLen(len(s)))
	n, err := decode(dst, []byte(s))
	if err != nil {
		return nil, fmt.Errorf("Failed to decode input (%s): %v",
			s, err)
	}
	if !verifyChecksum(dst[0:n]) {
		return nil, fmt.Errorf("Checksum mismatch.")
	}
	return dst[0 : n-kChecksumLength], nil
}
