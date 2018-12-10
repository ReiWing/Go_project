package main

import "fmt"

func main() {
	var price = 100
	fmt.Println("Price is", price, "dollars.")

	var taxRate = 0.08
	var tax = float64(price) * taxRate
	fmt.Println("Tax is", tax, "dollars.")

	var total = float64(price) + tax
	fmt.Println("Total cost is", total, "dollars.")

	var availableFunts = 120
	fmt.Println(availableFunts, "dollars available.0")
	fmt.Println("Within budget?", int(total) < availableFunts)
}
