package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup //// modified version of time.Sleep ; these are pointers

func main() {
	go greeter("Hello") // just fire up a thread which will be responsible to greet hello but we never waited for the thread to comeback and greet hello
	greeter("World")
	websitelist := []string{
		"http://go.dev",
		"http://react.dev",
		"http://fb.com",
		"http://gh.com",
	}

	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()
}

func greeter(s string) {
	for i := 0; i < 6; i++ {
		// solution
		time.Sleep(3 * time.Microsecond)
		fmt.Println(s)
	}
}

func getStatusCode(endpoint string) {
	defer wg.Done()

	result, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPS in endpoint")
	} else {
		fmt.Printf("%d 200 status code for website %s\n", result.StatusCode, endpoint)
	}
}
