package main

import "fmt"

func createPointer() *float64 {
	var myFloat = 98.5
	return &myFloat
}

func printPointer(myBoolPointer *bool) {
	fmt.Println(myBoolPointer)
}

func main() {
	var myInt int
	var myIntPointer *int
	myInt = 42
	myIntPointer = &myInt
	fmt.Println(*myIntPointer)

	var myFloatPointer *float64 = createPointer()
	fmt.Println(myFloatPointer, *myFloatPointer)

	var myBool bool = true
	printPointer(&myBool)
}
