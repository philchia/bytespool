package bytespool

import (
	"sync"
)

// SyncPool is a sync.Pool base slab allocation memory pool
type SyncPool struct {
	classes     []sync.Pool
	classesSize []int
	minSize     int
	maxSize     int
}

// New create a new bytes pool
func New(minSize, maxSize, factor int) *SyncPool {
	n := 0
	for chunkSize := minSize; chunkSize <= maxSize; chunkSize *= factor {
		n++
	}
	pool := &SyncPool{
		classes:     make([]sync.Pool, n),
		classesSize: make([]int, n),
		minSize:     minSize,
		maxSize:     maxSize,
	}
	n = 0
	for chunkSize := minSize; chunkSize <= maxSize; chunkSize *= factor {
		pool.classesSize[n] = chunkSize
		pool.classes[n].New = func(size int) func() interface{} {
			return func() interface{} {
				buf := make([]int64, size)
				return &buf
			}
		}(chunkSize)
		n++
	}
	return pool
}

// Make get a bytes slice from pool or make a new one with len = size
func (pool *SyncPool) Make(length, capacity int) []int64 {
	if length > capacity {
		panic("make: length out of range")
	}
	if capacity <= pool.maxSize {
		for i := 0; i < len(pool.classesSize); i++ {
			if pool.classesSize[i] >= capacity {
				mem := pool.classes[i].Get().(*[]int64)
				return (*mem)[:length]
			}
		}
	}
	return make([]int64, length, capacity)
}

// Free put bytes slice into pool or do nothing
func (pool *SyncPool) Free(mem []int64) {
	if size := cap(mem); size <= pool.maxSize {
		for i := 0; i < len(pool.classesSize); i++ {
			if pool.classesSize[i] >= size {
				pool.classes[i].Put(&mem)
				return
			}
		}
	}
}
