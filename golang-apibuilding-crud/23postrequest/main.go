package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Making a POST request using JSON data in GO")
	PerformPostJsonRequest()

}

// The route only accepts JSON data
// Whenever we are making a POST request, we might want the data we post to be stored onto a database
// To send the data to backend, we use JSON
// so let us create a fake Json Request
// step 1 - to mention to which url we are making a request
func PerformPostJsonRequest() {
	const myurl = "http://localhost:3000/post"

	// step 2 - to send data to the url

	requestBody := strings.NewReader(`
		{
			"coursename":"Lets go with Golang",
			"price": 0,
			"platform": "golang.co"
		}
	`)
	// step 3 - sending request
	response, err := http.Post(myurl, "application/json", requestBody) // ---> keep note of this "application/json"

	if err != nil {
		panic(err)
	}
	// step 4 - close the request
	defer response.Body.Close()
	// step 5 - reading the request
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}
