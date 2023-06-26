package database

import (
	"fmt"
	"github.com/Edouard127/FurryProtectorGo/utils"
	"runtime"
	"testing"
)

var cache = NewInMemoryCache[int32, int32](2000, false, 10000)

func TestInMemoryCache(t *testing.T) {
	for i := 0; i < 10000; i++ {
		cache.Set(int32(i), int32(i))
	}
	printMemStats()
}

func printMemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %s", utils.FormatRam(m.Alloc))
	fmt.Printf("\tTotalAlloc = %s", utils.FormatRam(m.TotalAlloc))
	fmt.Printf("\tSys = %s", utils.FormatRam(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}
