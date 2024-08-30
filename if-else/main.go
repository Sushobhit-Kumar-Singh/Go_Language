package main

import "fmt"

func main() {
	x := 21
	if x > 5 {
		fmt.Println("X is greater than 5")
	} else {
		fmt.Println("X is smaller than 5")
	}

	z := 10
	if z > 10 {
		fmt.Println("Z is greater than 10")
	} else if z > 5 {
		fmt.Println("Z is greater than 5 but smaller than 10")
	} else {
		fmt.Println("Z is smaller than 5")
	}

	y := 10
	if x > 5 && y > 5 { //if x > 5 && (y > 5 || z < 43)
		fmt.Println("Hey how are you ?")
	} else {
		fmt.Println("Lear programming with Hello World")
	}
}
