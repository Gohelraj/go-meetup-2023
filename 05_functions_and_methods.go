package main

import "fmt"

// Define a function that adds two numbers and returns the result
func add(a, b int) int {
	return a + b
}

// Define a Rectangle struct
type Rectangle struct {
	Length int
	Width  int
}

// Define a method for Rectangle that calculates the area
func (r Rectangle) Area() int {
	return r.Length * r.Width
}

func main() {
	// Call the add function and print the result
	result := add(2, 3)
	fmt.Println(result)

	// Create an instance of Rectangle
	rect := Rectangle{
		Length: 5,
		Width:  3,
	}

	// Call the Area method on the Rectangle instance and print the result
	area := rect.Area()
	fmt.Println(area)
}
