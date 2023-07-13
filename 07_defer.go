package main

import "fmt"

func main() {
	fmt.Println("Start of the program")

	// Defer statements are executed in LIFO (last-in, first-out) order
	defer fmt.Println("Deferred statement 1")
	defer fmt.Println("Deferred statement 2")

	fmt.Println("Middle of the program")

	// Defer can be used to ensure cleanup or closing of resources
	file := openFile("example.txt")
	defer closeFile(file)

	fmt.Println("End of the program")
}

// openFile simulates opening a file and returns a file handle
func openFile(filename string) string {
	fmt.Printf("Opening file: %s\n", filename)
	return "file-handle"
}

// closeFile simulates closing a file
func closeFile(file string) {
	fmt.Printf("Closing file: %s\n", file)
}
