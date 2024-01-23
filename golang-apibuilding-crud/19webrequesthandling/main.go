// Handling web requests in go
// net/http is the most fastest and recommended package used to handle web requests in golang
// It is always important to close the response in the end

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "http://golang.co" // step 1 - defining the url of the website

func main() {
	fmt.Println("Welcome to Web Request Handling in Go")

	// step 2 - making a http request

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is of type: %T\n", response)

	defer response.Body.Close() // step 3 -  It is caller's responsibility to close the connection

	// to read this response
	databytes, err := ioutil.ReadAll(response.Body) // ReadAll - What to read? --> response.Body

	if err != nil {
		panic(err)
	}
	// fmt.Println(string(databytes))
	content := string(databytes)
	fmt.Println(content)
}

// o/p ---> *http.Response - is a pointer. So we get the reference of the original response
// So we arent getting the copy of the response but getting the original response.
