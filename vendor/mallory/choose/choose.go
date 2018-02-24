package choose

import (
	c "crypto/rand"
	"encoding/binary"
	"io"
)

var randomIndex uint
var randomNumbers [100]uint64

type deadBeefReader struct{}

var deadBeef = [...]byte{0xde, 0xad, 0xbe, 0xef}

func (zr deadBeefReader) Read(buf []byte) (n int, err error) {
	for i := range buf {
		buf[i] = deadBeef[i%len(deadBeef)]
	}
	return len(buf), nil
}

// Pretend this is math/rand.New(rand.NewSource(99))
var fallbackReader = deadBeefReader{}

var _ io.Reader = (*deadBeefReader)(nil)

func fillRandomNumbers(reader *io.Reader, fallbackReader io.Reader) {
	var buf [8]byte
	for i := 0; i < len(randomNumbers); i++ {
		_, err := io.ReadAtLeast(*reader, buf[:7], 8)
		if err != nil {
			*reader = fallbackReader
		}
		randomNumbers[i] = binary.LittleEndian.Uint64(buf[:])
	}
}

func init() {
	fillRandomNumbers(&c.Reader, fallbackReader)
}

func ChooseString(strings []string) string {
	s := strings[int(randomNumbers[randomIndex%uint(len(randomNumbers))]%uint64(len(strings)))]
	randomIndex++
	return s
}
