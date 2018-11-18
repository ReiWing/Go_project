// pass_fail reports whether a grade is passing or failings
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, error := reader.ReadString('\n')
	if error != nil {
		log.Fatal(error)
	}

	input = strings.TrimSpace(input)
	grade, error := strconv.ParseFloat(input, 64)
	if error != nil {
		log.Fatal(error)
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
