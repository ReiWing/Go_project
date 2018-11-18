package main

import "fmt"

func double(number float64) float64 {
	return number * 2
}
func main() {
	dozen := double(22.333)
	fmt.Println(dozen)
	fmt.Println(double(0.222))
}
