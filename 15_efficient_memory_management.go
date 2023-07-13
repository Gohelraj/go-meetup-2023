package main

import (
	"fmt"
	"runtime"
	"strings"
)

func main() {
	PrintMemUsage()

	// Efficient string concatenation using strings.Builder
	var builder strings.Builder
	for i := 0; i < 10000; i++ {
		builder.WriteString("test")
	}

	text := builder.String()

	fmt.Println("Length of text:", len(text))
	PrintMemUsage() // Check memory usage after the concatenation operation
}

// PrintMemUsage prints the current memory usage stats. It prints values of:
// - Alloc: Current memory use
// - TotalAlloc: Total memory allocated over the lifetime of the program (may be higher than Alloc because of garbage collection)
// - Sys: The total memory obtained from the OS
// - NumGC: The number of times the garbage collector has been run
// For more details, please visit: https://golang.org/pkg/runtime/#MemStats
func PrintMemUsage() {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	fmt.Printf("Alloc = %v MiB", memStats.Alloc/1024/1024)
	fmt.Printf("\tTotalAlloc = %v MiB", memStats.TotalAlloc/1024/1024)
	fmt.Printf("\tSys = %v MiB", memStats.Sys/1024/1024)
	fmt.Printf("\tNumGC = %v\n", memStats.NumGC)
}
