package main

import (
	"fmt"
)

func recoverFromPanic() {
	// The recover function is called to capture the panic value.
	// If a panic occurred, it returns the value passed to panic.
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic:", r)
	}
}

func divide(x, y int) int {
	// The defer keyword is used to schedule the execution of the recoverFromPanic function
	// until the surrounding function, divide, returns.
	defer recoverFromPanic()

	if y == 0 {
		// If the divisor is zero, a panic is triggered with the message "division by zero".
		panic("division by zero")
	}

	return x / y
}

func main() {
	// The divide function is called with a valid divisor (2), and the result is printed.
	result := divide(10, 2)
	fmt.Println("Result:", result)

	// The divide function is called with a divisor of 0, triggering a panic.
	// The program will immediately stop execution and jump to the nearest deferred function, recoverFromPanic.
	result = divide(10, 0)
	fmt.Println("Result:", result)

	// After recovering from the panic, the program continues executing normally.
	fmt.Println("Program continues...")
}
