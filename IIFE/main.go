package main

import "fmt"

// A higher-order function that takes a function as an argument
func applyOperation(a, b int, operation func(int, int) int) int {
 return operation(a, b)
}

// Define specific operations
func add(x, y int) int {
 return x + y
}

func multiply(x, y int) int {
 return x * y
}

func main() {
 // Use higher-order function with different operations
 fmt.Println(applyOperation(3, 5, add))      // Output: 8
 fmt.Println(applyOperation(3, 5, multiply)) // Output: 15
}