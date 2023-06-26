package database

import (
	"fmt"
	"github.com/Edouard127/FurryProtectorGo/utils"
	"runtime"
	"testing"
)

func TestInMemoryCache(t *testing.T) {
	cache := NewInMemoryCache[byte, byte](2000, false, 255)
	for i := 0; i < 255; i++ {
		cache.Set(byte(i), byte(i))
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
