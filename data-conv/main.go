package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int = 42
	fmt.Println("Number is ", num)
	fmt.Printf("Type of num is %T\n", num)

	// num = num + 1.23 error because int me float value add kr rhe h

	var data float64 = float64(num)
	data = data + 1.23
	fmt.Println("Data is ", data)
	fmt.Printf("Type of data is %T\n", data) //only type change hota na ki data ki value

	num = 123
	str := strconv.Itoa(num)
	fmt.Println("str is ", str)
	fmt.Printf("Type of str is %T\n", str)

	number_string := "1234"
	//number_int = number_int + 2576 //eroor because data mismatch string+ int
	number_int, _ := strconv.Atoi(number_string)
	//number_int = number_int + 2576 //correct
	fmt.Println("Number_int is ", number_int)
	fmt.Printf("Type of number_int is %T\n", number_int)

	num_string := "3.14"
	number_float, _ := strconv.ParseFloat(num_string, 64)
	fmt.Println("Number_float is ", number_float)
	fmt.Printf("Type of number_float is %T\n", number_float)

}
