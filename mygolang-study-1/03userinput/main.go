package main

import (
	"bufio"
	"fmt"
	"os"
)

// to take input from the user (example: googleforms takes input from user and the code just contains the logic)

func main() {
	welcome := "Welcome to user input program"
	fmt.Println(welcome)

	// Reader is reading the i/p from the terminal
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza:")

	// comma ok syntax(err err) - whatever the reader reads, I need to store that in a variable caled input and if there is any error it will be stored in err

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("The type of rating is: %T", input)

}
