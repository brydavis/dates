package main

import (
	"fmt"
	"log"
)

func main() {
	// t1 := "7/21/1980 20:07:34.830" // Original datetime is non-standard (NS) format
	t1 := "21/07/80 10:07:34.830AM" // Original datetime is standard (S) format

	code, err := Validate(t1)
	if err != nil {
		log.Fatal(err)
	}

	format, _ := FindFormat(code)
	fmt.Printf("Your datetime matches the standard format (%s)\n\r", format)

	badCode := 99
	_, err = FindFormat(badCode)
	if err != nil {
		fmt.Print(err)
	}

	newCode := 110
	newFormat, _ := FindFormat(newCode)
	fmt.Printf("Updating datetime to new standard format: code (%d) => format (%s)\n\r", newCode, newFormat)
	t2 := Convert(t1, format, newCode)
	fmt.Println(t2) // 1980-07-21
}
