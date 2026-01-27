package main

import "fmt"

func hello() {
	fmt.Println("Hello, World!")
}

func add(a int, b int) {
	fmt.Println("Adding", a, "and", b)
	fmt.Println("Sum is", a+b)
}

func plus(a int, b int) int {
	reult := a + b
	return reult
}

func message() string {
	return "This is a message from the message function."
}

func main() {
	hello() // Call the function in main function
	fmt.Println("-----------------------------------------------------------------")
	add(3, 4) // Call the add function with arguments
	fmt.Println("-----------------------------------------------------------------")
	sum := plus(5, 6) // Call the plus function and store the return value
	fmt.Println("Returned sum is", sum)
	fmt.Println("-----------------------------------------------------------------")
	fmt.Println(message()) // Call the message function and print the return value
	fmt.Println("-----------------------------------------------------------------")
}
