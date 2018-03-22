package bytespool

import (
	"testing"
)

func BenchmarkPoolMake(b *testing.B) {
	pool := New(4, 1<<20, 2)
	b.SetParallelism(10)
	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			bts := pool.Make(100, 100)
			bts[0] = 8
			handle(bts)
			pool.Free(bts)
		}
	})
}

func BenchmarkMake(b *testing.B) {
	b.SetParallelism(10)
	b.ResetTimer()
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			bts := make([]byte, 100, 100)
			bts[0] = 8
			handle(bts)
		}
	})
}

func handle(bts []byte) {
	l := len(bts)
	cap := cap(bts)
	if bts[0] == 8 {
		bts[0] = 1
	}
	_ = l
	_ = cap
}
