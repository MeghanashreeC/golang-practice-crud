// how to use files - we will need ioutils and defer packages

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")

	// aim - we have some content and we need to put them into a text file
	content := "This needs to go in a file - meghana.in"

	// working with the text file
	file, err := os.Create("./myltogofile.txt") // os.Create to create a file
	checkNilErr(err)
	//if err != nil {
	//	panic(err)
	// } // common way of handling errors

	// to write something
	length, err := io.WriteString(file, content) //writestring calculates the total number of characters and writes to a variable
	checkNilErr(err)
	fmt.Println("length is:", length)
	defer file.Close() // must close the file in the end. It is always recommended to use defer
	readFile("./myltogofile.txt")

}

// To read a file -- ioutil
func readFile(filename string) { // read the entire file and returns the number as bytes
	databyte, err := ioutil.ReadFile(filename) // the filename parameter in the readFile function is essentially the path to the file you want to read. The function then reads the content of the specified file and prints it to the console.
	checkNilErr(err)
	fmt.Println("Text data inside the file is:\n", string(databyte)) // using string to print the line as it is and not in numbers

} // utility to read a file is ioutil -- > ioutil can possibly throw an error

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
