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
func greet(name string string {
	fmt.Printf("Hello, %s!\n", name)
}

func main() {
	// Variable declared but not used
	var unusedVar int

	// Logical error: passing a string to factorial
	result := factorial("5")
	fmt.Println("Factorial of 5 is:", result)

	// Calling a function with the wrong number of arguments
	sum := addNumbers(3)
	fmt.Println("Sum is:", sum)

	// Attempting to access an out-of-range index in a slice
	slice := []int{1, 2, 3}
	fmt.Println("Fourth element:", slice[3])

	// Infinite loop (logical error)
	for {
		fmt.Println("Infinite loop")
	}
}

// Function definition mismatch (expects two arguments but the body has three)
func addNumbers(a int, b int) int {
	return a + b + c
}

// Function missing a return statement
func divide(a int, b int) int {
	if b == 0 {
		fmt.Println("Error: Division by zero")
	}
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
	fmt.Println(i.(int))
}

// Deliberate misspelling of "fmt"
func typoExample() {
	fmtx.Println("This won't work")
}

// Struct with missing fields in initialization
func structExample() {
	type Person struct {
		Name string
		Age  int
	}

	p := Person{
		Name: "John",
	}
	fmt.Println(p)
}

// Map without proper initialization
func mapExample() {
	var m map[string]int
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
