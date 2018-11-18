package main

import (
	"fmt"
	"log"
)

func paintNeeded(width float64, height float64) (float64, error) {
	if width <= 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid.\n", width)
	} else if height <= 0 {
		return 0, fmt.Errorf("a height of %0.2f is invalid.\n", height)
	}
	area := width * height
	return area / 10.0, nil
}

func main() {

	// err := errors.New("the height should be positive")
	// fmt.Println(err.Error())
	// log.Fatal(err)
	// fmt.Printf("a height of %0.2f is invalid", -2.3333)

	var amount, total float64
	var err error
	amount, err = paintNeeded(4.2, -2.22222)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%0.2f liters needed.\n", amount)
	}
	total += amount
	amount, err = paintNeeded(4.2, -2.22222)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%0.2f liters needed.\n", amount)
	}
	total += amount
	fmt.Printf("Total: %0.2f liters needed.\n", total)
}
