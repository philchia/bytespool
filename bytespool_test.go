package bytespool

import (
	"fmt"
	"testing"
)

func TestAlloc(t *testing.T) {
	pool := New(4, 1<<20, 2)

	bts := pool.Make(34, 34)
	defer pool.Free(bts)

	fmt.Println(len(bts), cap(bts))
	for _, b := range bts {
		fmt.Print(b)
	}
}
