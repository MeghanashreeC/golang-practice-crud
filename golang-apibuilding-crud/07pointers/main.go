package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers")

	// var ptr *int // * indicates that the datatype is a pointer. So this ptr here is responsible for storing the integer
	// fmt.Println("Value of pointer is", ptr)

	myNumber := 23

	var ptr = &myNumber                              // & is used as reference to something. So here & is referencing to the memoru address of myNumber
	fmt.Println("Value of actual pointer is:", ptr)  // gives us actual memory address
	fmt.Println("Value of actual pointer is:", *ptr) // gives us the value of myNumber

	*ptr = *ptr * 2 // * takes the actual value of ptr(myNumber)
	// & takes the address of the ptr)(myNumber)

	fmt.Println("New value is:", *ptr)
}
