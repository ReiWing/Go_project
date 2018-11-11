package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileInfo, error := os.Stat("README.txt")
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println(fileInfo.Size(), " bytes")
}
