package main

import (
	"fmt"
	"practice/golang-apibuilding-crud/testing/calculator"
)

func main() {
	fmt.Println("Calculator!!")

	result := calculator.Add(3, 4)
	fmt.Println(result)
}
