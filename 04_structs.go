package main

import "fmt"

// Define a struct type called "Person"
type Person struct {
	Name string
	Age  int
}

func main() {
	// Create a new instance of the Person struct
	person := Person{
		Name: "Raj Gohel",
		Age:  27,
	}

	// Access and modify the fields of the struct
	fmt.Println("Person Name:", person.Name)
	fmt.Println("Person Age:", person.Age)

	// Update the fields of the struct
	person.Name = "Jane Smith"
	person.Age = 35

	fmt.Println("Updated Person Name:", person.Name)
	fmt.Println("Updated Person Age:", person.Age)
}
