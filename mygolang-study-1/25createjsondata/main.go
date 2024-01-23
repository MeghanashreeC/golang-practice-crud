package main

import (
	"encoding/json"
	"fmt"
)

// creating a struct

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`        // to make it common
	Password string   `json:"-"`              // to remove password
	Tags     []string `json:"tags,omitempty"` // if the value is nil, dont throw the field at all (no space b/w tags,omitempty)
}

func main() {
	fmt.Println("A little more on JSON")
	EncodeJSON()
	DecodeJson()

}

// We will be working with encoding the JSON. I want to convert all the data to valid JSON

func EncodeJSON() { // Encoding means converting the data into valid JSON
	// create a slice
	lcoCourses := []course{
		{"ReactJS", 299, "golang.co", "abc123", []string{"web-dev", "JS"}}, // adding values to struct
		{"JS", 399, "golang.co", "sdhdc123", []string{"abc-dev", "PS"}},
		{"Golang", 499, "golang.co", "hfhc123", nil},
	}

	// Let us package this data as JSON data

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t") //created by library JSON  // Marshal is an implementation of JSON
	// \t helps to indent the data. it is always best to use MarshalIndent
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)

}

// how to consume json data --> how to take json data from web
func DecodeJson() {
	jsonDataFromWeb := []byte(`    
	
	{
		"coursename": "ReactJS",
		"Price": 299,
		"website": "golang.co",
		"tags": ["web-dev","JS"]
	}
	`)
	// whenever data comes from the web, it is going to be in byte format

	// to check if the json data is valid
	var lcoCourse course
	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON is valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse) // to pass the reference use &
		fmt.Printf("%#v\n", lcoCourse)

	} else {
		fmt.Println("JSON IS NOT VALID")
	}

	// there are some cases where we just want to add data to key-value pair
	// If i dont want to create structure and I want to extract data from key value pair

	var myOnlineData map[string]interface{} // since the data is coming from online web, I am not sure if the rest of the part is always a string. It might be arrays also. So using interfaces
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	// for loop

	for k, val := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and type of data is %T\n", k, val, val)
	}

}
