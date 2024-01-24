// channels are a way in which multiple goroutines can talk to each other.

package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in golang")

	// creating a channel -

	myCh := make(chan int, 2)

	// create goroutines
	wg := &sync.WaitGroup{}

	// myCh <- 5
	// fmt.Println(<-myCh)
	wg.Add(2)
	go func(ch <-chan int, wg *sync.WaitGroup) {

		val, IsChanelOpen := <-myCh
		fmt.Println(IsChanelOpen)
		fmt.Println(val)
		// fmt.Println(<-myCh) // receiving channel value  for 5
		wg.Done()
	}(myCh, wg)

	go func(ch chan<- int, wg *sync.WaitGroup) {
		myCh <- 0 // adding channel value
		close(myCh)
		// myCh <- 6
		wg.Done()
	}(myCh, wg)

	wg.Wait()

}
