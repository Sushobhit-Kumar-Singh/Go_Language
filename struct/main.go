package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Contact struct {
	Email string
	Phone string
}

type Address struct {
	Home  int
	Area  string
	State string
}

type Employee struct {
	Person_Details Person
	Person_Contact Contact
	Person_Address Address
}

func main() {
	var shobhit Person
	fmt.Println("Shobhit Person :", shobhit) //""""0//__0 print hua h

	shobhit.FirstName = "Shobhit"
	shobhit.LastName = "Singh"
	shobhit.Age = 25
	fmt.Println("Shobhit Person :", shobhit)
	// 2nd method
	person1 := Person{
		FirstName: "Akash",
		LastName:  "Kumar",
		Age:       26,
	}
	fmt.Println("Person 1 :", person1)

	// new keyword
	var person2 = new(Person) //new keyword is a pointer
	person2.FirstName = "Prince"
	person2.LastName = "Agarwal"
	person2.Age = 24

	fmt.Println("Person 2 :", person2)
	fmt.Println("Person 2 :", person2.FirstName)
	fmt.Println("Age of Shobhit is ", shobhit.Age)

	var employee1 Employee
	employee1.Person_Details = Person{
		FirstName: "Bittu",
		LastName:  "Singh",
		Age:       21,
	}
	employee1.Person_Contact.Email = "contact@gmail.com"
	employee1.Person_Contact.Phone = "9931114422"

	employee1.Person_Address = Address{
		Home:  12,
		Area:  "Patna",
		State: "Bihar",
	}
	fmt.Println("Employee 1 : ", employee1)
	fmt.Println("Employee 1 :", employee1.Person_Contact.Email)

}
