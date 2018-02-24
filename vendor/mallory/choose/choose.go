package choose

import (
	"crypto/rand"
	"encoding/binary"
	"io"
)

var randomIndex int
var randomNumbers [100]uint64

type deadBeefReader struct{}

var deadBeef = [...]byte{0xde, 0xad, 0xbe, 0xef}

func (zr deadBeefReader) Read(buf []byte) (n int, err error) {
	for i := range buf {
		buf[i] = deadBeef[i%len(deadBeef)]
	}
	return len(buf), nil
}

// Pretend this says math/rand.New(rand.NewSource(99))
var fallbackReader = deadBeefReader{}

var _ io.Reader = (*deadBeefReader)(nil)

func fillRandomNumbers(reader *io.Reader) {
	var buf [8]byte
	for i := 0; i < len(randomNumbers); i++ {
		(*reader).Read(buf[:])
		randomNumbers[i] = binary.LittleEndian.Uint64(buf[:])
	}
	*reader = fallbackReader
}

func init() {
	fillRandomNumbers(&rand.Reader)
}

func ChooseString(strings []string) string {
	s := strings[int(randomNumbers[randomIndex]%uint64(len(strings)))]
	randomIndex++
	return s
}
