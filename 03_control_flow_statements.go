package main

import "fmt"

func main() {
	// If-Else statement
	x := 10
	if x > 5 {
		fmt.Println("x is greater than 5")
	} else {
		fmt.Println("x is less than or equal to 5")
	}

	// For loop with range
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// Iterate over numbers 1 to 5 and check if each number is even or odd
	for i := 1; i <= 5; i++ {
		if i%2 == 0 {
			fmt.Println(i, "is even")
		} else {
			fmt.Println(i, "is odd")
		}
	}

	// Use a switch statement to identify the selected language
	language := "Go"
	switch language {
	case "Python":
		fmt.Println("Selected language is Python")
	case "Java":
		fmt.Println("Selected language is Java")
	case "Go":
		fmt.Println("Selected language is Go")
	default:
		fmt.Println("Unknown language")
	}

	// For is Go's while
	num := 5
	// Loop continues as long as num is greater than 0
	for num > 0 {
		fmt.Println(num)
		num--
	}
}
