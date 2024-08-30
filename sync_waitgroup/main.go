package main

import (
	"fmt"
	"sync"
)

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done() // signal that this goroutine is done
	fmt.Printf("worker %d started\n", i)
	// some task is happining
	fmt.Printf("worker %d end\n", i)

}

func main() {
	// fmt.Println("Explore gorouyine started")

	var wg sync.WaitGroup
	// start 3 worker goroutine
	for i := 1; i <= 3; i++ {
		wg.Add(1) // increament the waitgroup counter
		// worker(i)
		// go worker(i)
		go worker(i, &wg)

	}
	// wait for all workers to finish
	wg.Wait()
	fmt.Println("workers task completed")
}
