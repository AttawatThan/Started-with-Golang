package main

import "fmt"

func main() {
	// Call Slice
	var courseName []string
	courseName = []string{"Rust", "Python", "Golang", "JavaScript"}
	fmt.Println(courseName)
	courseName = append(courseName, "TypeScript", "C#") // append new elements to slice
	fmt.Println(courseName)

	//Select elements from slice
	courseDS := courseName[:2] // [Include, Exclude)]
	fmt.Println(courseDS)
}
