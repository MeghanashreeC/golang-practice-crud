package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to timestudy in golang")

	// To print present time

	presentTime := time.Now()
	fmt.Println(presentTime)
	// To change the format of the time

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	// to create the time from some values

	createdDate := time.Date(2020, time.August, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))

}

// go gives us tools to create exe(executable files ---> 'go env')
// GOOS="linux" go build  to create exe files

// use "go mod tidy" to install all the packages in the module
