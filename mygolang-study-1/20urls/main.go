// Inorder to grab some information from the url like parameters

package main

import (
	"fmt"
	"net/url"
)

// create a fictious url

const myurl string = "http://golang.com:3000/learn?coursename=reactjs&paymentid=fghjjhgf23"

func main() {
	fmt.Println("Welcome to handling urls in Go")
	fmt.Println(myurl)

	// parsing url - to extract some information

	result, err := url.Parse(myurl) // url.Parse ---> method
	if err != nil {
		panic(err)
	}
	fmt.Println(result.Scheme) // result contains so much of info and the first info is Scheme
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Port()) // it is a port

	// to get query parameters individually --> result.RawQuery --> we can store all that is there here into a variable
	// to store result.RawQuery into a variable qparams

	qparams := result.Query()                                 // .Query stores the parameter into a better format
	fmt.Printf("The type of query params are: %T\n", qparams) // url.Values - key-value pair

	fmt.Println(qparams["coursename"]) // key value pair  --> to get the value

	for _, val := range qparams {
		fmt.Println("Params are:", val)
	}

	// to construct a url from a chunk of information (reverse procedure)

	partsofUrl := &url.URL{ // use & always --> &url.URL
		Scheme:  "https",
		Host:    "golang.com",
		Path:    "/learn",
		RawPath: "user=meghana",
	}

	// construct a url

	anotherURL := partsofUrl.String()
	fmt.Println(anotherURL)

}
