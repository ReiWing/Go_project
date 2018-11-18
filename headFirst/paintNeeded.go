package main

import (
	"fmt"
)

func paintNeeded(width float64, heigth float64) float64 {
	area := width * heigth
	return area / 10.0
}

func main() {

	// err := errors.New("the height should be positive")
	// fmt.Println(err.Error())
	// log.Fatal(err)
	// fmt.Printf("a height of %0.2f is invalid", -2.3333)

	var amount, total float64
	amount = paintNeeded(4.2, 5.6)
	fmt.Printf("%0.2f liters needed.\n", amount)
	total += amount
	amount = paintNeeded(5.2, 22.3)
	fmt.Printf("%0.2f liters needed.\n", amount)
	total += amount
	fmt.Printf("%0.2f liters needed.\n", total)
}
