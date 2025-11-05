package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Welcome to Go Coverage Exploration!")

	// Call some functions
	result := Add(5, 3)
	fmt.Printf("5 + 3 = %d\n", result)

	greeting := Greet("World")
	fmt.Println(greeting)
}

// Add returns the sum of two integers
// This function will be FULLY tested
func Add(a, b int) int {
	return a + b
}

// Subtract returns the difference of two integers
// This function will be FULLY tested
func Subtract(a, b int) int {
	return a - b
}

// Square returns the square of an integer
// This function is currently untested
func Square(n int) int {
	return n * n
}

// Multiply returns the product of two integers
// This function will be PARTIALLY tested (only one branch)
func Multiply(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}

	result := a * b

	// This branch will NOT be covered
	if result < 0 {
		fmt.Println("Result is negative!")
	}

	return result
}

// Divide performs division with error handling
// This function will be PARTIALLY tested (only success case)
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

// Greet returns a greeting message
// This function will be PARTIALLY tested (only some branches)
func Greet(name string) string {
	if name == "" {
		return "Hello, stranger!"
	}

	if strings.ToLower(name) == "world" {
		return "Hello, World!"
	}

	// This will be tested
	if len(name) > 10 {
		return fmt.Sprintf("Hello, %s! That's a long name!", name)
	}

	// This will NOT be tested
	return fmt.Sprintf("Hello, %s!", name)
}

// IsEven checks if a number is even
// This function will be COMPLETELY UNTESTED
func IsEven(n int) bool {
	return n%2 == 0
}

// IsPositive checks if a number is positive
// This function will be COMPLETELY UNTESTED
func IsPositive(n int) bool {
	if n > 0 {
		return true
	}
	return false
}

// ComplexFunction demonstrates multiple branches
// This will have VERY PARTIAL coverage
func ComplexFunction(x int, y int, z bool) string {
	var result string

	// This branch will be tested
	if x > y {
		result = "x is greater"
		if z {
			result += " and z is true"
		}
	} else if x < y {
		// This will NOT be tested
		result = "y is greater"
		if z {
			result += " and z is true"
		} else {
			result += " and z is false"
		}
	} else {
		// This will NOT be tested
		result = "x and y are equal"
	}

	return result
}
