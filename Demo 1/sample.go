package main

import (
	"fmt"
	"strings" // Deliberately not used to show unused import
)

// This function calculates the factorial of a number
func factorial(n int) int {
	if n < 0 {
		fmt.Println("Error: Negative numbers are not allowed")
		return -1
	}
	if n == 0 {
		return 1
	}
	return n * factorial(n - 1)
}

// Deliberate syntax error (missing closing parenthesis)
func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

func main() {
	// Variable declared but not used
	var unusedVar int

	// Logical error: passing a string to factorial
	result := factorial(5)
	fmt.Println("Factorial of 5 is:", result)

	// Calling a function with the wrong number of arguments
	sum := addNumbers(3, 4) // Fixed function call to pass two arguments
	fmt.Println("Sum is:", sum)

	// Correcting the out-of-range index access
	slice := []int{1, 2, 3}
	if len(slice) > 3 {
		fmt.Println("Fourth element:", slice[3])
	} else {
		fmt.Println("Index out of range. The slice only has", len(slice), "elements.")
	}

	// Infinite loop (logical error)
	for {
		fmt.Println("Infinite loop")
	}
}

// Function definition mismatch (expects two arguments but the body has three)
func addNumbers(a int, b int) int {
	return a + b // Fixed function body to return a + b
}

// Function missing a return statement
func divide(a int, b int) int {
	if b == 0 {
		fmt.Println("Error: Division by zero")
		return 0 // Fixed by returning a value
	}
	return a / b
}

// Deliberate naming conflict
func conflictFunction() {
	fmt.Println("Conflict 1")
}

func conflictFunction() {
	fmt.Println("Conflict 2")
}

// Incorrect type assertion
func typeAssertionExample() {
	var i interface{} = "hello"
	// Fixed type assertion to avoid panic
	if str, ok := i.(string); ok {
		fmt.Println(str)
	} else {
		fmt.Println("Type assertion failed")
	}
}

// Deliberate misspelling of "fmt"
func typoExample() {
	// Fixed the typo
	fmt.Println("This works now")
}

// Struct with missing fields in initialization
func structExample() {
	type Person struct {
		Name string
		Age  int
	}

	// Fixed the struct initialization to include both fields
	p := Person{
		Name: "John",
		Age:  30,
	}
	fmt.Println(p)
}

// Map without proper initialization
func mapExample() {
	// Fixed map initialization
	m := make(map[string]int)
	m["key"] = 42
	fmt.Println(m)
}

// Unused constants
const Pi = 3.14
const Greeting = "Hello"

// Function that never gets called
func unusedFunction() {
	fmt.Println("This function is never called")
}

