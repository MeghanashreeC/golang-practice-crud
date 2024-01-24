package main

import "fmt"

func main() {
	fmt.Println("Mutex in Golang")
}

// Mutex
// It provides a lock over the memory
// It am going to lock this memory until one goroutine is working and till the time it is writing it will not allow anybody to use this memory.

// Read-Write Mutex

// GOROUTINES
// Goroutines are useful when you want to do multiple things simultaneously.
// For example, if you have ten things you want to do at the same time, you can do each one on a separate goroutine, and wait for all of them to finish.
// Note: You should not be putting context. Background() in your functions
