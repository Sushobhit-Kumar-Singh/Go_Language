package main

import "fmt"

func main() {

	numbers := []int{1, 2, 3, 4, 5}
	numbers = append(numbers, 3, 12, 10, 12, 15)
	fmt.Println("Number :", numbers)
	fmt.Printf("Number has data type : %T\n", numbers)
	fmt.Println("Length :", len(numbers))

	name := []string{}
	fmt.Println("Name : ", name)

	num := []int{1, 2, 3}
	fmt.Println("Slice:", num)
	fmt.Println("Length:", len(num))
	fmt.Println("Capacity:", cap(num))

	number := make([]int, 3, 5) //3-length,5-capacity
	number = append(number, 4)
	number = append(number, 98)
	number = append(number, 6) //capacity double ho jayega 5*2=10 because len 5 se jayada ho gya isliye
	fmt.Println("Slice:", number)
	fmt.Println("Length:", len(number))
	fmt.Println("Capacity:", cap(number))

}
