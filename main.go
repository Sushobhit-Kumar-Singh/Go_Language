package main

import (
	"fmt"
	//"mylearning/myutil"
)

func main() {
	fmt.Println("Learning Go Language by Hello World")
	//myutil.PrintMessage("Hello World")
	var name string = "Shobhit" //76 error I define string type so it is  not correct for 76 but "76"-correct
	var version = "version Latest"
	var versions = 76 // I dont define any type so it is correct
	fmt.Println(name)
	fmt.Println(version)
	fmt.Println(versions)

	var money int = 67000
	var b, c int = 1, 2
	fmt.Println(b, c)
	var currency = 3489
	fmt.Println(money)
	fmt.Println(currency)

	var dimension float64 = 87.12
	fmt.Println(dimension)

	var decide bool = false
	fmt.Println(decide)

	var person = "Kumar" //23
	fmt.Println(person)

	const pi = 3.14
	// pi = 23 not change becaue of const
	fmt.Println(pi)

	persons := 123 //"shobhit kumar" without var ka use kiye bina variable decler kiya its shortcut
	fmt.Println(persons)

	var Public = "Data is important"
	var private = "data is private"
	fmt.Println(Public)
	fmt.Println(private)

}
