package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	// fmt.Println("Starting of the program")
	// sum := add(5, 6)
	// fmt.Println("Sum is : ", sum)
	// defer fmt.Println("Middle of the program") //defer keyword use krne se program main function khatam hone ke bad run hota h
	// fmt.Println("End of the program")

	// agar 2 ke aage defer use hua h tb  o stack ko follo krta h
	// yani 2nd defer wala pahle print hoga
	fmt.Println("Starting of the program")
	data := add(5, 6)
	defer fmt.Println("Data is : ", data)
	defer fmt.Println("Middle of the program")
	fmt.Println("End of the program")

}
