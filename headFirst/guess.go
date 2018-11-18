// guess challenges players to guess a random number

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
	succes := false
	for guesses := 0; guesses < 10; guesses++ {

		fmt.Println("You have ", 10-guesses, " guesses left")
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
			fmt.Println("Good job! You guessed it!")
			succes = true
			break
		}
	}
	if succes != true {
		fmt.Println("Sorry you didn't guess my number. It was ", target, ".")
	}
}
