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

	// Inter the number of walls
	fmt.Print("Enter the number of walls: ")
	reader_walls := bufio.NewReader(os.Stdin)
	wallsCount, error := reader_walls.ReadString('\n')
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
		reader_width := bufio.NewReader(os.Stdin)
		width_temp, error := reader_width.ReadString('\n')
		if error != nil {
			log.Fatal(error)
		}
		width_temp = strings.TrimSpace(width_temp)
		width, error := strconv.ParseFloat(width_temp, 64)
		if error != nil {
			log.Fatal(error)
		}

		//Enter the height of wall
		fmt.Print("Enter the height of wall number ", x, ": ")
		reader_height := bufio.NewReader(os.Stdin)
		height_temp, error := reader_height.ReadString('\n')
		if error != nil {
			log.Fatal(error)
		}
		height_temp = strings.TrimSpace(height_temp)
		height, error := strconv.ParseFloat(height_temp, 64)
		if error != nil {
			log.Fatal(error)
		}

		paintNeeded(width, height)
	}
}

func paintNeeded(width, height float64) {
	fmt.Printf("%0.2f liters needed.\n", width*height/10.0)
}
