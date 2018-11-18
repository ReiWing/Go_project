package main

import (
	"fmt"
	"math"
)

func squareRoot(number float64) (float64, error) {
	if number < 0 {
		return 0, fmt.Errorf("can't get square root of negative number")
	} else {
		return math.Sqrt(number), nil
	}
}

func main() {
	root, err := squareRoot(-8)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%0.3f\n", root)
	}
}
