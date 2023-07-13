package main

import (
	"fmt"
	"time"
)

// simulate a time-consuming task
func processTask(taskID int) {
	time.Sleep(2 * time.Second)
	fmt.Println("Task", taskID, "completed")
}

func main() {
	start := time.Now()

	// Execute tasks one by one
	for i := 1; i <= 5; i++ {
		processTask(i)
	}

	elapsed := time.Since(start)
	fmt.Println("Total time:", elapsed)
}
