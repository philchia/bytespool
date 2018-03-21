package bytespool

import (
	"fmt"
	"testing"
)

func TestAlloc(t *testing.T) {
	pool := New(4, 1<<20, 2)

	bts := pool.Alloc(34)
	defer pool.Free(bts)

	fmt.Println(len(bts), cap(bts))
	for _, b := range bts {
		fmt.Print(b)
	}
	fmt.Println()
	m := make([]byte, 5, 34)
	fmt.Println(len(m), cap(m))
	for _, b := range m {
		fmt.Print(b)
	}
	fmt.Println()

	m = m[:32]
	for _, b := range m {
		fmt.Print(b)
	}
}
