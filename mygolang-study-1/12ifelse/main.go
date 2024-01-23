package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to ifelse in golang")

	loginCount := 30
	var result string

	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		result = "Exactly 10"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("The number is even")

	} else {
		fmt.Println("The number is odd")
	}

	if num := 3; num < 10 {
		fmt.Println("The number is less than 10")

	} else {
		fmt.Println("The number is greater than 10")
	}

	//if err != nil {

	// }
}
