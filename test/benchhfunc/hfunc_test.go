package benchhfunc

import (
	"crypto/aes"
	"github.com/cespare/xxhash"
	"testing"
)

func BenchmarkAes(b *testing.B) {
	block, err := aes.NewCipher([]byte("1234567890123456"))
	if err != nil {
		panic(err)
	}
	result := make([]byte, aes.BlockSize)
	src := []byte("hello world")
	if len(src)%aes.BlockSize != 0 {
		paddingBlock := make([]byte, aes.BlockSize-(len(src)%aes.BlockSize))
		src = append(src, paddingBlock...)
	}
	for i := 0; i < b.N; i++ {
		block.Encrypt(result, src)
	}
}

func BenchmarkXXHash(b *testing.B) {
	src := []byte("hello world")
	for i := 0; i < b.N; i++ {
		xxhash.Sum64(src)
	}
}
