package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("Numbers is :", i)
	}

	counter := 0
	for {
		fmt.Println("Infinite Loop")
		counter++

		if counter == 3 {
			break
		}
	}

	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("Index of Number is %d, and value is %d\n", index, value)
	}

	data := "Hello, World!"
	for index, char := range data {
		fmt.Printf("Index of data is %d, and value is %c\n", index, char)
	}
}
