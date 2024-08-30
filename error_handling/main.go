package main

import "fmt"

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("denominator must not be zero")
	}
	return a / b, nil
}

func main() {
	fmt.Println("started error handling in function")
	ans, err := divide(10, 0)
	//ans, _ := divide(10, 0) fuction return kr rhe h per hum ignore krna ho

	if err != nil {
		fmt.Println("Error handaling")
	}
	fmt.Println("Division of two numbers is ", ans)
}
