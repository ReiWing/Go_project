// pass_fail reports whether a grade is passing or failings
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, error := reader.ReadString('\n')
	/*
		if score == 100 {
			fmt.Println("Flawless victory!")
		} else if score >= 60 {
			fmt.Println("Go ahead!")
		} else {
			fmt.Println("You shell not pass!")
		}
	*/
	if error != nil {
		log.Fatal(error)
	}
	fmt.Print(input)
}
