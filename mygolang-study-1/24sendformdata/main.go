package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println("Sending form data in Golang")
	PerformPostFormRequest()
}

// We need to have form data in cases like image upload,etc

func PerformPostFormRequest() {
	const myurl = "http://localhost:3000/postform"

	// create a form data

	data := url.Values{}             // it is going to be data and that data is coming from url -- > url.Values is a key value pair
	data.Add("firstname", "meghana") // adding key and values
	data.Add("lastname", "shree")
	data.Add("email", "meghana@co.co")

	// creating a request

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	// create content

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}
