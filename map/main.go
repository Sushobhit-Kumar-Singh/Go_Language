package main

import "fmt"

func main() {
	// name <-> grade
	studentGrades := make(map[string]int)

	studentGrades["Shobhit"] = 100
	studentGrades["Alice"] = 34
	studentGrades["Avinash"] = 95
	studentGrades["Bob"] = 85

	fmt.Println("Marks of Bob :", studentGrades["Bob"])
	studentGrades["Bob"] = 100 // for update
	fmt.Println("Marks of Bob :", studentGrades["Bob"])

	delete(studentGrades, "Bob") //for delete
	fmt.Println("Marks of Bob :", studentGrades["Bob"])

	//checking if a key exists
	grades, exists := studentGrades["David"]
	fmt.Println("Grades of David :", grades)
	fmt.Println("David exists: ", exists)

	fmt.Println("Marks of David :", studentGrades["David"]) //checking if a key exists

	//checking if a key exists
	Grades, Exists := studentGrades["Shobhit"]
	fmt.Println("Grades of Shobhit :", Grades)
	fmt.Println("Shobhit exists: ", Exists)

	for index, value := range studentGrades {
		fmt.Printf("Key is %s and marks is %d\n", index, value)
	}

	person := map[string]int{
		"Alice":   90,
		"Bob":     85,
		"Charlin": 95,
	}
	for index, value := range person {
		fmt.Printf("------Key is %s and marks is %d\n", index, value)
	}

}
