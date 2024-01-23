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
	fmt.Printf("Name is %v and email is %v\n", meghana.Name, meghana.Email)
	meghana.GetStatus() // method
	meghana.GetDetails()
	meghana.NewMail()
	fmt.Printf("Name is %v and email is %v\n", meghana.Name, meghana.Email) // to view the changes print again
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

// defining a func

func (u User) GetStatus() { // creating a func(alias Struct name) methodname(){}
	fmt.Println("Is user active:", u.Status)
}

func (u User) GetDetails() {
	fmt.Println("The details of user are:", u.Name, u.Age, u.Email)
}

// we can manipulate based on method

func (u User) NewMail() {
	u.Email = "test@go.com"
	fmt.Println("New Email of user is:", u.Email)

}
