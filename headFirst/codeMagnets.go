package main

import "fmt"

func main() {

	var originalCount, eatenCount int = 10, 4

	fmt.Print("I started with ", originalCount, " apples. ")
	fmt.Print("Some jerk ate ", eatenCount, " apples. ")
	fmt.Println("There are", originalCount-eatenCount, "apples left.")
}
