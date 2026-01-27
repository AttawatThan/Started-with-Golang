package main

import "fmt"

var product = make(map[string]float64) //map[key type]value type

func main() {
	fmt.Println("product = ", product)

	// Adding key-value pairs to the map
	product["Laptop"] = 1500.00
	product["Smartphone"] = 800.00
	product["Tablet"] = 400.00
	fmt.Println("Updated product = ", product)

	// Delete a key-value pair from the map
	delete(product, "Tablet") // Deleting the key "Tablet" with product variable
	fmt.Println("After deletion, product = ", product)

	// Updating a value in the map
	product["Laptop"] = 888.88
	fmt.Println("After updating, product = ", product)

	// Accessing a value from the map
	value1 := product["Smartphone"]
	fmt.Println("Price of Smartphone =", value1)

	courseName := map[string]string{
		"101": "Rust",
		"102": "Python",
		"103": "Golang",
	}
	fmt.Println("courseName =", courseName)
}
