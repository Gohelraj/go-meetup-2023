package main

import "fmt"

func main() {
	// Declare a variable
	number := 42
	fmt.Println("Original value:", number)

	// Pass the variable to a function that modifies it using a pointer
	updateValue(&number)
	fmt.Println("Updated value:", number)

	// Create a new variable of type int pointer
	var pointer *int
	pointer = &number
	fmt.Println("Pointer value:", *pointer)

	// Modify the value indirectly through the pointer
	*pointer = 99
	fmt.Println("Modified value:", number)
}

// Function that updates the value of a variable through a pointer
func updateValue(value *int) {
	*value = 100
}
