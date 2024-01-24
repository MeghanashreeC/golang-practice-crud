package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race Condition - Golang")
	// wait group

	wg := &sync.WaitGroup{}

	mut := &sync.RWMutex{} //RWMutex --> ReadWriteMutex

	// Lets create a couple of goroutines
	var score = []int{0}

	// Eventually I want to push more values into score and thats why I will be using goroutines
	wg.Add(3)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("One R")
		mut.Lock() // very important to lock after performing write operation
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	// wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Two R")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	// wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Three R")
		mut.RLock()
		score = append(score, 3)
		mut.RUnlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println(score)

}
