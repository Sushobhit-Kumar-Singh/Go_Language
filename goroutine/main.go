package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello, World!")
	time.Sleep(2000 * time.Millisecond) // Simulating some work
	fmt.Println("sayHello function end successfully")
}

func sayHi() {
	fmt.Println("Hi Shobhit ")
	time.Sleep(1000 * time.Millisecond) // Simulating some work

}

func main() {
	fmt.Println("Learning Goroutines")
	// sayHello()
	go sayHello()
	// sayHi()
	go sayHi()

	// wait for a moment for to allow the goroutine
	time.Sleep(1000 * time.Millisecond)
	// time.Sleep(3000 * time.Millisecond)

}
