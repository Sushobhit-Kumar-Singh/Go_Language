package main

import (
	"fmt"
	"os"
)

func main() {

	// for create file
	/* file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error while creating file: ", err)
		return
	}
	defer file.Close()

	content := "hello world by Shobhit"
	byte, error := io.WriteString(file, content+"\n")
	fmt.Println("Byte written :", byte)
	//_, error := io.WriteString(file, content+"\n") //byte ke jagah _ use kiye h , error // +"\n"--> for cursur new line me jane ke liye example file m
	if error != nil {
		fmt.Println("Error while creating file: ", error)
		return
	}

	fmt.Println("Successfully created file")
	*/

	/*
		file, err := os.Open("example.txt")
		if err != nil {
			fmt.Println("Error while opening file: ", err)
			return
		}
		defer file.Close()

		// Creater a buffer to read the file content
		buffer := make([]byte, 1024)

		// Read the file content into the buffer
		for {
			n, err := file.Read(buffer)
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println("Error while reading file", err)
				return
			}

			// Process the read content
			fmt.Println(string(buffer)[:n])
		}
	*/
	// Read the entire file into a byte slice
	// content, err := ioutil.ReadFile("example.txt") // old version
	content, err := os.ReadFile("example.txt") // new version both are same

	if err != nil {
		fmt.Println("Error while reading file ", err)
		return
	}
	fmt.Print(string(content))

}
