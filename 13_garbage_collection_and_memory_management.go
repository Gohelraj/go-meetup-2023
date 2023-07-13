package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// Variable to store our 2D slices
	var memoryHolder [][]int

	for index := 0; index < 4; index++ {

		// Here we create a new slice with a large capacity (999999).
		// This slice is stored inside our memoryHolder variable to prevent it from being garbage collected.
		// The aim is to continually increase our program's memory footprint
		// so we can observe Go's memory management though our PrintMemUsage function.
		slice := make([]int, 0, 999999)
		memoryHolder = append(memoryHolder, slice)

		// We print memory statistics after each memory allocation
		PrintMemUsage()
		// Then we wait for a second
		time.Sleep(time.Second)
	}

	// Here we clear our memoryHolder by setting it to nil and freeing up the memory
	memoryHolder = nil
	// We print the memory statistics after the clearing
	PrintMemUsage()

	// We then force the Garbage Collector to run, this should reduce the memory usage significantly
	runtime.GC()
	// We print the memory stats after forcing the garbage collection
	PrintMemUsage()
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
