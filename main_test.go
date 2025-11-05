package main

import (
	"testing"
)

// FULLY TESTED FUNCTIONS

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 5, 3, 8},
		{"negative numbers", -5, -3, -8},
		{"mixed signs", -5, 3, -2},
		{"with zero", 0, 5, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 5, 3, 2},
		{"negative result", 3, 5, -2},
		{"with zero", 5, 0, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Subtract(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Subtract(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// PARTIALLY TESTED FUNCTIONS

func TestMultiply(t *testing.T) {
	// Only testing positive numbers - negative branch won't be covered
	result := Multiply(5, 3)
	if result != 15 {
		t.Errorf("Multiply(5, 3) = %d; want 15", result)
	}

	// Test zero case
	result = Multiply(0, 5)
	if result != 0 {
		t.Errorf("Multiply(0, 5) = %d; want 0", result)
	}

	// NOTE: Not testing negative results, so the "result < 0" branch is uncovered
}

func TestDivide(t *testing.T) {
	// Only testing successful division - not testing the error case
	result, err := Divide(10, 2)
	if err != nil {
		t.Errorf("Divide(10, 2) unexpected error: %v", err)
	}
	if result != 5 {
		t.Errorf("Divide(10, 2) = %d; want 5", result)
	}

	// NOTE: Not testing divide by zero, so error handling is uncovered
}

func TestGreet(t *testing.T) {
	// Test empty string
	result := Greet("")
	if result != "Hello, stranger!" {
		t.Errorf("Greet(\"\") = %q; want %q", result, "Hello, stranger!")
	}

	// Test "world" case
	result = Greet("World")
	if result != "Hello, World!" {
		t.Errorf("Greet(\"World\") = %q; want %q", result, "Hello, World!")
	}

	// Test long name
	result = Greet("VeryLongName")
	expected := "Hello, VeryLongName! That's a long name!"
	if result != expected {
		t.Errorf("Greet(\"VeryLongName\") = %q; want %q", result, expected)
	}

	// NOTE: Not testing regular short names, so default branch is uncovered
}

func TestComplexFunction(t *testing.T) {
	// Only testing one branch
	result := ComplexFunction(10, 5, true)
	expected := "x is greater and z is true"
	if result != expected {
		t.Errorf("ComplexFunction(10, 5, true) = %q; want %q", result, expected)
	}

	// Test with z false
	result = ComplexFunction(10, 5, false)
	expected = "x is greater"
	if result != expected {
		t.Errorf("ComplexFunction(10, 5, false) = %q; want %q", result, expected)
	}

	// NOTE: Not testing x < y or x == y branches
}

// COMPLETELY UNTESTED FUNCTIONS
// IsEven - no tests
// IsPositive - no tests
