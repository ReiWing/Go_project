package main

import "fmt"

func main() {
	quantity := 4
	length, width := 1.2, 2.4
	customerName := "Rei Wing"

	fmt.Print(customerName)
	fmt.Print(" has ordered ", quantity, " sheets")
	fmt.Print(" each with an area of ")
	fmt.Println(length*width, "square meters.")

}
