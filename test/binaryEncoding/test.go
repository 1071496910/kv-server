package main

import (
	"encoding/binary"
	"fmt"
	"github.com/cespare/xxhash"
	"unsafe"
)

func main() {
	fmt.Println(binary.Uvarint([]byte("12345678")))

	buf := make([]byte, 8)
	binary.PutUvarint(buf, 12345)

	a := []byte{1, 0, 0, 0, 0, 0, 0, 0}
	fmt.Println(*(*uint)(unsafe.Pointer(&a[0])))

	h := xxhash.New()
	h.Write([]byte("hello world"))
	fmt.Println(string(h.Sum(nil)))

}
