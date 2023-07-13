package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// simulate a time-consuming task
func processTask(taskID int, wg *sync.WaitGroup, errCh chan<- error) {
	defer wg.Done()

	// Simulating an error condition for task 3
	if taskID == 3 {
		errCh <- errors.New("error occurred in task 3")
		return
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Task", taskID, "completed")
}

func main() {
	start := time.Now()

	// Create a WaitGroup to wait for all tasks to complete
	var wg sync.WaitGroup

	// Create a channel to receive error messages
	errCh := make(chan error)

	// Number of tasks to execute
	numTasks := 5

	// Add the number of tasks to the WaitGroup
	wg.Add(numTasks)

	// Execute tasks concurrently
	for i := 1; i <= numTasks; i++ {
		go processTask(i, &wg, errCh)
	}

	// Start a Goroutine to handle errors from the error channel
	go func() {
		for err := range errCh {
			fmt.Println("Error:", err)
		}
	}()

	// Wait for all tasks to complete
	wg.Wait()

	// Close the error channel
	close(errCh)

	elapsed := time.Since(start)
	fmt.Println("Total time:", elapsed)
}
