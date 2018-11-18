package main

import (
	"fmt"
)

func main() {
	var width, height, area float64
	width = 4.2
	height = 3.0
	area = width * height
	fmt.Printf("%0.2f liters needed.\n", area/10.0)

	width = 5.2
	height = 3.5
	area = width * height
	fmt.Printf("%0.2f liters needed.\n", area/10.0)

	fmt.Printf("The %s cost %d cents each.\n", "gumballs", 23)
	fmt.Printf("That will be $ %0.2f please.\n", 0.23*5)

	fmt.Printf("A float: %f\n", 3.1415)
	fmt.Printf("An integer: %d\n", 15)
	fmt.Printf("A string: %s\n", "Go to sleep already!")
	fmt.Printf("A boolean: %t\n", false)
	fmt.Printf("Values: %v %v %v\n", 1.2, "\t", true)
	fmt.Printf("Values: %#v %#v %#v\n", 1.2, "\t", true)
	fmt.Printf("Types: %T %T %T\n", 1.2, "\t", true)
	fmt.Printf("Percent sign: %%\n")

	fmt.Printf("%0.2f liters needed.\n", 1.8199999999999998)

	fmt.Printf("%12s | %s\n", "Product", "Cost in cents")
	fmt.Println("---------------------------------")
	fmt.Printf("%12s | %2d\n", "Stamps", 50)
	fmt.Printf("%12s | %2d\n", "Paper Clips", 5)
	fmt.Printf("%12s | %2d\n", "Tape", 99)
}
