package main

import "fmt"

func main() {
	fmt.Println(" We are learning Golang")

	var name [5]string //golang me array dynamic type ka nhi hota h iske liye hun Slice use krte h
	name[0] = "Shobhit"
	name[2] = "Avinash"

	fmt.Println("Names of person is :", name)

	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Numbers is :", numbers)

	fmt.Println("Length of Numbers array is :", len(numbers))
	fmt.Println("values of name at 2 index is :", name[2])

	var price [5]string
	price[2] = "Prince"
	price[0] = "Akash"
	fmt.Println("Price is :", price)
	fmt.Printf("Price Array is %q\n", price)

}
