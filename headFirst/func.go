package main

import "fmt"

func functionA(a int, b int) {
	fmt.Println(a + b)
}

func functionB(a int, b int) {
	fmt.Println(a * b)
}

func funtionC(a bool) {
	fmt.Println(!a)
}

func functionD(a string, b int) {
	for i := 0; i < b; i++ {
		fmt.Println(a)
	}
	fmt.Println()
}

func passFail(grade float64) string {
	if grade < 60 {
		return "failing"
	}
	return "passing"
}

func main() {

	functionA(2, 3)
	functionB(2, 3)
	funtionC(true)
	functionD("$", 4)
	functionA(5, 6)
	functionB(5, 6)
	funtionC(false)
	functionD("ha", 3)
	fmt.Println(passFail(60.2))
	fmt.Println(passFail(33.1))
}
