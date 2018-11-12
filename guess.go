package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("I've choosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Make a guess: ")
	input, error := reader.ReadString('\n')
	if error != nil {
		log.Fatal(error)
	}

	input = strings.TrimSpace(input)
	guess, error := strconv.Atoi(input)
	if error != nil {
		log.Fatal(error)
	}

	if guess < target {
		fmt.Println("Oops. Your guess was LOW.")
	} else if guess > target {
		fmt.Println("Oops. Your gusss was HIGH.")
	} else {
		fmt.Println("You are great!")
	}

	fmt.Println(target)
}
