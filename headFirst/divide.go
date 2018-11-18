package main

import (
	"errors"
	"fmt"
)

func divide(divident float64, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, errors.New("Can't divide by zero")
	}
	return divident / divisor, nil
}

func main() {
	quotient, err := divide(22.123, 0.2222222222222)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%0.2f\n", quotient)
	}
}
