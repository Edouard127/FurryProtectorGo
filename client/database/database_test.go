package database

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestInMemoryCache(t *testing.T) {
	cache := NewInMemoryCache[int32, int32](2000, 10000)
	printMemStats()
	for i := 0; i < 10000; i++ {
		cache.Set(int32(i), int32(i))
	}
	printMemStats()
	simulateWait(1000)
	printMemStats()
	for i := 0; i < 10000; i++ {
		cache.Get(int32(i))
	}
	simulateWait(4000)
	fmt.Println(cache.Get(4694))
	cache.Close()
	printMemStats()
}

func simulateWait(timeout int) {
	time.Sleep(time.Duration(timeout) * time.Millisecond)
}

func printMemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
