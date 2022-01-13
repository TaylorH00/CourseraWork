package main

import (
	"fmt"
	"strings"
)

func main() {
	var userInput string

	fmt.Print("Enter a string:")

	fmt.Scan(&userInput)

	userInput = strings.ToLower(userInput)

	if strings.HasPrefix(userInput, "i") && strings.HasSuffix(userInput, "n") && strings.Contains(userInput, "a") {
		fmt.Print("Found!")
	} else {
		fmt.Print("Not Found!")
	}
}
