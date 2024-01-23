package main

import "fmt"

func main() {
	fmt.Println("Welcome to maps")

	languages := make(map[string]string)

	languages["JS"] = "JavaScript" // JS = key and JavaScript = value
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages:", languages)
	fmt.Println("JS shorts for:", languages["JS"])

	// TO DELETE
	delete(languages, "RB")
	fmt.Println("List of all languages:", languages)

	// LOOPS

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	} // For key = Js, the value will be JavaScript

}
