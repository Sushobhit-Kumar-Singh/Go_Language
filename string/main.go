package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "apple,orange,banana"     // agar .(dot) hota tb
	parts := strings.Split(data, ",") //, ke jagah .(dot) likhte
	fmt.Println(parts)

	str := "one two three four two two five"
	count := strings.Count(str, "two")
	fmt.Println("Count: ", count)

	strs := "    Hello, Go !     " // "   Hello     go" bich ka space trim nhi hota h
	trimmed := strings.TrimSpace(strs)
	fmt.Println("Trimmed: ", trimmed)

	str1 := "Shobhit"
	str2 := "Singh"
	//result := strings.Join([]string{str1, str2}, " ")
	result := strings.Join([]string{str1, "Kumar", str2}, " ")

	fmt.Println("Result: ", result)
}
