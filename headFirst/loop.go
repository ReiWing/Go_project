package main

import "fmt"

func main() {
	for x := 1; x <= 3; x++ {
		fmt.Print(x)
	}

	for x := 3; x >= 1; x-- {
		fmt.Print(x)
	}

	for x := 2; x <= 3; x++ {
		fmt.Print(x)
	}

	for x := 1; x < 3; x++ {
		fmt.Print(x)
	}

	for x := 1; x <= 3; x += 2 {
		fmt.Print(x)
	}

	for x := 1; x >= 3; x++ {
		fmt.Print(x)
	}

}
