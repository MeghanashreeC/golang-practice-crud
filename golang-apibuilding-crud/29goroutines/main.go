package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var signals = []string{"test"}

var wg sync.WaitGroup //// modified version of time.Sleep ; these are pointers

// Mutex
var mut sync.Mutex // pointer --> because we need to point it to multiple goroutines

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
	fmt.Println(signals)
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

		// mutex lock and unlock
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d 200 status code for website %s\n", result.StatusCode, endpoint)
	}
}
