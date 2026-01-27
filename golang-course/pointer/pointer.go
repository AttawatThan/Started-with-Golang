package main

import "fmt"

func zeroValue(ivalue int) {
	ivalue = 0

}

func zeroPointer(ipointer *int) {
	*ipointer = 0 // dereference the pointer and set the value to 0
}

func main() {
	i := 1
	fmt.Println("i = ", i)

	zeroValue(i)
	fmt.Println("i from function zeroValue = ", i) // i is still 1

	zeroPointer(&i)                                        // & is mean pass the address of i
	fmt.Println("i value from function zeroPointer = ", i) // i is now 0
	fmt.Println("i address from function zeroPointer = ", &i)

}
