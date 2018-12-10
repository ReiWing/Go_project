// Paint needed calculator for couple of walls

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	var totalArea float64
	totalArea = 0

	fmt.Print("Liter per liters: ")
	readerLiters := bufio.NewReader(os.Stdin)
	litersTemp, error := readerLiters.ReadString('\n')
	if error != nil {
		log.Fatal(error)
	}
	litersTemp = strings.TrimSpace(litersTemp)
	liters, error := strconv.ParseFloat(litersTemp, 64)
	if error != nil {
		log.Fatal(error)
	}

	// Inter the number of walls
	fmt.Print("Enter the number of walls: ")
	readerWalls := bufio.NewReader(os.Stdin)
	wallsCount, error := readerWalls.ReadString('\n')
	if error != nil {
		log.Fatal(error)
	}
	wallsCount = strings.TrimSpace(wallsCount)
	number, error := strconv.Atoi(wallsCount)
	if error != nil {
		log.Fatal(error)
	}

	// Loop for 'Enter the wall's stats' flow
	for x := 1; x <= number; x++ {

		//Enter the width of wall
		fmt.Print("Enter the width of wall number ", x, ": ")
		readerWidth := bufio.NewReader(os.Stdin)
		widthTemp, error := readerWidth.ReadString('\n')
		if error != nil {
			log.Fatal(error)
		}
		widthTemp = strings.TrimSpace(widthTemp)
		width, error := strconv.ParseFloat(widthTemp, 64)
		if error != nil {
			log.Fatal(error)
		}

		//Enter the height of wall
		fmt.Print("Enter the height of wall number ", x, ": ")
		readerHeight := bufio.NewReader(os.Stdin)
		heightTemp, error := readerHeight.ReadString('\n')
		if error != nil {
			log.Fatal(error)
		}
		heightTemp = strings.TrimSpace(heightTemp)
		height, error := strconv.ParseFloat(heightTemp, 64)
		if error != nil {
			log.Fatal(error)
		}

		area := width * height
		fmt.Printf("%0.2f liters needed for this wall.\n", area/liters)
		totalArea = totalArea + area
		fmt.Printf("Total area now is: %0.2f\n", totalArea)
	}
	fmt.Printf("%0.2f liters needed for all walls.\n", totalArea/liters)
}

func paintNeeded(width, height float64) {
	area := width * height / 10.0
	fmt.Printf("%0.2f liters needed.\n", area)
}
