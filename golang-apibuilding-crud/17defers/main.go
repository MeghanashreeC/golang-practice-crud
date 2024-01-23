// defer - When a func executes, it executes line by line even if it compiled.When we use
// defer keyword, the line after this keyword is not executed but it executes the last line of the function(reverse order)
// LIFO - Last in First Out

package main

import "fmt"

// func main() {
// defer fmt.Println("World")     // prints last line first and then prints defer
// fmt.Println("Hello")
// }

func main() {
	defer fmt.Println("One")          //fifth
	defer fmt.Println("Two")          // fourth
	defer fmt.Println("Three")        // third
	fmt.Println("My name is Meghana") // first
	myDefer()                         // second

}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i) // 43210 --> o/p
	}

}
