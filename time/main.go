package main

import (
	"fmt"
	"time"
)

func main() {

	currentTime := time.Now()
	fmt.Println("Current time: ", currentTime)
	fmt.Printf("Type of currentTime %T\n", currentTime)

	//formatted := currentTime.Format("02-01-2006, Monday") //error dd-mm-yyyy
	//formatted := currentTime.Format("2006/01/02, 15:04:05")// 24 hours format
	formatted := currentTime.Format("2006/01/02, 03:04:05 PM") // for am/pm 12 hours format

	fmt.Println("Formatted Time: ", formatted)

	layout_str := "2006-01-02" //"02/01/2006"
	dateStr := "2024-08-20"    //"25/08/2024"
	formatted_time, _ := time.Parse(layout_str, dateStr)
	fmt.Println("Formatted time:", formatted_time)

	// add 1 more day in currentTime
	new_date := currentTime.Add(24 * time.Hour)
	fmt.Println("New_Date : ", new_date)
	formatted_new_date := new_date.Format("2006/01/02 Monday")
	fmt.Println("Formatted_new_date time: ", formatted_new_date)
}
