package main

import (
	"fmt"
	"log"

	"github.com/Grade-calc/gradeinput"
)

func main() {
	fmt.Print("Enter a grade")
	grade, err := gradeinput.InputGrade()
	if err != nil {
		log.Fatal(err)
	}

	var status string
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Println("A grade of", grade, "is", status)
}
