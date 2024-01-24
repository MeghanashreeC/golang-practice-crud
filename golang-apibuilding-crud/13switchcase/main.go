package main

import (
	"fmt"
	"log"
	"math/big"
	"math/rand"
)

func main() {
	fmt.Println("Welcome to Switch case in Go")

	// rand.Seed(time.Now().UnixNano()) // whenever there is rand you must Seed it
	// diceNumber := rand.Intn(6) + 1
	// fmt.Println("Value of dice is: ", diceNumber)

	MyRandomNum, err := rand.Int(rand.Reader, big.NewInt(5)) // rand.Reader comes from crypto package only
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(MyRandomNum)

	// switch

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
		fallthrough
	case 2:
		fmt.Println("You can move two spots")
		fallthrough
	case 3:
		fmt.Println("You can move three spots")
	case 4:
		fmt.Println("You can move four spots")
	case 5:
		fmt.Println("You can move five spots")
	case 6:
		fmt.Println("You can move six spots and role dice again")
	default:
		fmt.Println("What was that!")
	}

}

// if fallthrough is mentioned, the next statement will also be printed
