package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza app")
	fmt.Println("Please rate our pizza between 1 and 5")

	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println("Thanks for rating, ", input)

	// If I want to add 1 to whatever rating the user is giving

	// numRating = input + 1   // error because input is of type string and we cannot add 1
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64) // to handle the above error i.e., strconv helps in conversion

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating: ", numRating+1)
	}

}

// strings.TrimSpace(input), 64 ---> used to trim off the \n to prevent error
