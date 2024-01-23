package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in Go")
	// declaring a slice
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Friday"}
	fmt.Println(days)

	// first loop - for

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	// second type

	for i := range days {
		fmt.Println(days[i])
	}

	// third type - for each

	for index, day := range days {
		fmt.Printf("index is %v and value is %v\n", index, day)
	}

	// others // break and continue

	rougueValue := 1

	for rougueValue < 10 {

		if rougueValue == 2 {
			goto lco
		}

		if rougueValue == 5 {
			break // loop breaks after reaching 5
			// break -- stops and terminates after 4
			// continue -- skips only 5 and prints the others
		}
		fmt.Println("Value is:", rougueValue)
		rougueValue++
	}

	// goto statements

lco:
	fmt.Println("Jumping at LearnCode.in")

	// see up for goto

}
