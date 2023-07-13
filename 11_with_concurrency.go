package main

import (
	"fmt"
	"time"
)

// simulate a time-consuming task
func processTask(taskID int, ch chan<- bool) {
	time.Sleep(2 * time.Second)
	fmt.Println("Task", taskID, "completed")
	ch <- true
}

func main() {
	start := time.Now()

	// Create a buffered channel with a capacity of 5
	ch := make(chan bool, 5)

	// Execute tasks concurrently
	for i := 1; i <= 5; i++ {
		go processTask(i, ch)
	}

	// Wait for all tasks to complete
	for i := 1; i <= 5; i++ {
		<-ch
	}

	elapsed := time.Since(start)
	fmt.Println("Total time:", elapsed)
}
