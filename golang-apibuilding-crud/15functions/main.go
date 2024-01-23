package main

import "fmt"

func main() { // main acts as an entry point of the prgm to start the prgm
	fmt.Println("Welcome to functions in golang")
	greeter() // calling the below func here
	greeterTwo()

	proResult, myMessage := proAdder(2, 3, 4, 2)
	fmt.Println("Pro result is:", proResult)
	fmt.Println("Pro message is:", myMessage)

	// I want to create a special method to add two numbers

	result := adder(3, 5)
	fmt.Println("Result is:", result)

	// func greeterTwo() {
	//	fmt.Println("Another method")
	//}
	// greeterTwo()   // we cannot have a func within another func
}

func adder(val1 int, val2 int) int { // function signature - what kind of value we pass and what kind of value we expect to be returned
	return val1 + val2
}

// I have no idea how many values are coming but I want to add all of them(pro functions)
func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}

	return total, "Hi Pro result"
}

func greeterTwo() {
	fmt.Println("Another Method")

}

// creating a simple function
func greeter() {
	fmt.Println("Hello from golang")
} // we dont get this o/p because if we are declaring this function, we need to call this function as well
