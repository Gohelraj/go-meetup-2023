package main

import "fmt"

func main() {
	// Explicitly defined variable - type is specified explicitly
	var message string = "Hello from Go"
	
	// Constant variables
	const pi = 3.14159
	const maxAttempts = 5
	
	// Within a function, variables can be defined using shorthand syntax without using the var keyword
	// Shorthand syntax uses := instead of =
	message1 := "Hello World" // 'message1' is of type string

	// Multiple variables can be defined together
	var x, y int // both 'x' and 'y' are defined as int

	// Multiple variables can be defined together and initialized
	var x1, y1 int = 5, 10

	// Multiple variables can be defined together with initialization
	name, age, isDefined := "Go", 13, true

	// Shorthand variable declaration with implicit type inference and initialization
	slice := []int{1, 2, 3}

	// Short declaration with explicit type and initialization
	var array [3]int = [3]int{4, 5, 6}

	// Short declaration with explicit type and initialization
	var mapVar map[string]int = make(map[string]int)
	mapVar["one"] = 1
	mapVar["two"] = 2

	// Print statements
	fmt.Println("message =", message)
	fmt.Println("x =", x, "y =", y)
	fmt.Println("x1 =", x1, "y1 =", y1)
	fmt.Println("name =", name, "age =", age, "isDefined =", isDefined)
	fmt.Println("message1 =", message1)
	fmt.Println("slice =", slice)
	fmt.Println("array =", array)
	fmt.Println("mapVar =", mapVar)
	fmt.Println("pi =", pi)
	fmt.Println("maxAttempts =", maxAttempts)
}
