package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Welcome to web verb - get")
	PerformGetRequest()
}

// create a seperate method to perform get request --- 1

func PerformGetRequest() {
	const myurl = "http://localhost:3000/get"
	// make a get request  ---- 2
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	// remember to close   --- 3
	defer response.Body.Close()

	fmt.Println("Status code:", response.StatusCode)
	fmt.Println("Content length:", response.ContentLength)
	// to read this response  --> use ioutil  ---- 4
	content, _ := ioutil.ReadAll(response.Body) // the one that you use to close
	fmt.Println(string(content))
}
