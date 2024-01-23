package main

import (
	"fmt"
)

const LoginToken string = "uhgfcdvhcj" // L capital indicates public variable

func main() {
	var username string = "Meghana"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var IsLoggedIn bool = false
	fmt.Println(IsLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.345678
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default values and aliases

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// implicit type

	var website = "learn.co"
	fmt.Println(website) // Go takes default type if not mentioned

	// no var style

	numberOfUser := 300000.0
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)

}
