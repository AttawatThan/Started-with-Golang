package main

import "fmt"

var numberInt int = 42 // Package level Global variable, var can be used anywhere in the package
var numberInt2 int = 100
var msg string = "Hello, Go!"

func main() {
	numberFloat := 3.14 // Call Short Variable Declaration, := can only be used inside functions
	fmt.Println("Hello, World!")
	fmt.Println(numberInt)
	fmt.Println(numberFloat)
	fmt.Println(msg)

	// plust number inside Println
	fmt.Println("Sum with same type:", numberInt+numberInt2)                // Same type int
	fmt.Println("Sum with different type:", float64(numberInt)+numberFloat) // Different type int and float64, must convert to same type

	// Concat string
}
