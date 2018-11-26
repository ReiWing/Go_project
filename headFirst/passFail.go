// pass_fail reports whether a grade is passing or failings
package main

import (
	"fmt"
	"keyboard"
	"log"
)

func main() {
	fmt.Println("Enter grade: ")
	grade, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}

	var status string
	if grade > 100 {
		status = "Die cheater!"
	} else if grade == 100 {
		status = "Flawless victory!"
	} else if grade >= 60 {
		status = "Go ahead!"
	} else {
		status = "You shell not pass!"
	}

	fmt.Println("Accept your destiny...", status)
}
