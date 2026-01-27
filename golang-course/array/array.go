package main

import "fmt"

// Call Array
var productName [4]string   // Default value is empty string ""
var productPrice [4]float64 // Default value is 0.0

func main() {
	fmt.Println(productName)
	fmt.Println(productPrice)
	// Assign value to array
	productName[0] = "Macbook Pro"
	productName[1] = "Macbook Air"
	productName[2] = "Macbook Mini"
	productName[3] = "Macbook Studio"
	fmt.Println(productName[0])

	price := [4]float64{1999.99, 1499.99, 999.99, 2499.99} // Declare and assign value to array, shorterm format
	fmt.Println(price)
}
