// random function --> rand

package main

import (
	"fmt"
	"log"
	"math/big"

	// "math/rand"
	"crypto/rand"
)

func main() {
	fmt.Println("Maths in golang")

	// var mynumberOne int = 2
	// var mynumberTwo float64 = 4.5

	// fmt.Println("The sum is:", mynumberOne+int(mynumberTwo))

	// random number
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5) + 1)

	// random from crypto

	MyRandomNum, err := rand.Int(rand.Reader, big.NewInt(5)) // rand.Reader comes from crypto package only
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(MyRandomNum)

}
