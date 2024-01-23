package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("Welcome to slices")

	var fruitList = []string{"Apple", "Banana", "Mango"}
	fmt.Printf("Type of fruitlist is %T\n", fruitList)

	fruitList = append(fruitList, "Chickoo", "Orange")
	fmt.Println(fruitList)

	// To slice out the first element in fruitlist
	fruitList = append(fruitList[1:3]) // [Banana Mango] first element is no longer available
	fmt.Println(fruitList)

	fruitList = append(fruitList[0:3]) // [Apple, Banana, Mango]

	// utilising the make keyword in slices

	highScore := make([]int, 4) // syntax for make

	// default allocation is 4
	highScore[0] = 234
	highScore[1] = 945
	highScore[2] = 465
	highScore[3] = 867

	// we can deallocate it in go

	highScore = append(highScore, 555, 666, 321)

	fmt.Println(highScore)

	// sort package - it helps sort strings, int, float and also boolean

	fmt.Println(sort.IntsAreSorted(highScore)) // checks if highscore is sorted i.e., prints true or fals
	sort.Ints(highScore)                       // sorts the array
	fmt.Println(highScore)                     // prints the sorted array

	// how to remove a value from slice based on index

	var courses = []string{"Reactjs", "JavaScript", "Swift", "python", "Ruby"}
	fmt.Println(courses)

	// To remove swift
	var index int = 2
	courses = append(courses[0:index], courses[index+1:]...) // ... -> syntax
	fmt.Println(courses)

}
