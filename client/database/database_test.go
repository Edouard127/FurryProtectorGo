package database

import (
	"fmt"
	"github.com/Edouard127/FurryProtectorGo/utils"
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
	fmt.Printf("Alloc = %s", utils.FormatRam(m.Alloc))
	fmt.Printf("\tTotalAlloc = %s", utils.FormatRam(m.TotalAlloc))
	fmt.Printf("\tSys = %s", utils.FormatRam(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}
