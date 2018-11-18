package main

import (
	"fmt"
	"math"
)

func manyReturns() (int, bool, string) {
	return 1, true, "hello"
}

func floatParts(number float64) (integerPart int, fractionalPart float64) {
	wholeNumber := math.Floor(number)
	return int(wholeNumber), number - wholeNumber
}

func main() {
	myInt, myBool, myString := manyReturns()
	fmt.Println(myInt, myBool, myString)

	cans, reminder := floatParts(22.26)
	fmt.Printf(cans, reminder)
}
