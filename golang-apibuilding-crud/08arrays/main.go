package main

import "fmt"

func main() {

	fmt.Println("Welcome to arrays")

	// Method 1
	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[2] = "Mango"
	fruitList[3] = "Orange"

	fmt.Println("My array is:", fruitList)
	fmt.Println("Length of fruitlist is:", len(fruitList))

	// Method 2

	var vegList = [3]string{"potato", "tomato", "onion"}
	fmt.Println("My veggies are:", vegList)
	fmt.Println("Length of my vegList is:", len(vegList))
}
