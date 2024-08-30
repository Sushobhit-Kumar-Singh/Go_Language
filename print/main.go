package main

import "fmt"

func main() {
	age := 25
	name := "Alice"
	height := 5.8234567

	fmt.Println("age: ", age, "height: ", height, "name: ", name)
	fmt.Println("Hello World")

	fmt.Printf("age is %d\n", age)
	fmt.Printf("height is %.3f\n", height)
	fmt.Printf("name is %s\n", name)

	fmt.Printf("Name: %s, Age: %d, Height: %.2f\n", name, age, height)

	fmt.Printf("Type of name is %T\n", name) //data type janane ke liye T ka use krte h
	fmt.Printf("Type of age is %T\n", age)
	fmt.Printf("Type of height is %T\n", height)

}
