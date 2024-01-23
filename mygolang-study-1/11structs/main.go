package main

import "fmt"

func main() {
	fmt.Println("Welcome to structs")

	// no inheritance in golang; no super or parent or child concepts in golang
	meghana := User{"Meghana", "meghana@go.dev", true, 23}
	fmt.Println(meghana)

	// to read it easier with details
	// Whenever using structs, use "%+v"
	fmt.Printf("Meghana details are: %+v\n", meghana)

	// to see just one value
	fmt.Printf("Name is %v and email is %v", meghana.Name, meghana.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
