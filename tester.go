package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var userInput string
	fmt.Print("Enter a string: ")
	fmt.Scan(&userInput)

	intsUsed := make([]int, 0)

	for foundAll != true {
		switch {
		case strings.ContainsAny(userInput, "one"):
			intsUsed = append(intsUsed, 1)
		case strings.ContainsAny(userInput, "two"):
			intsUsed = append(intsUsed, 2)
		case strings.ContainsAny(userInput, "thre"):
			intsUsed = append(intsUsed, 3)
		case strings.ContainsAny(userInput, "four"):
			intsUsed = append(intsUsed, 4)
		case strings.ContainsAny(userInput, "five"):
			intsUsed = append(intsUsed, 5)
		case strings.ContainsAny(userInput, "six"):
			intsUsed = append(intsUsed, 6)
		case strings.ContainsAny(userInput, "sevn"):
			intsUsed = append(intsUsed, 7)
		case strings.ContainsAny(userInput, "eight"):
			intsUsed = append(intsUsed, 8)
		case strings.ContainsAny(userInput, "nine"):
			intsUsed = append(intsUsed, 9)
		default:
			foundAll = f
		}
	}

	intsUsed = sort.IntSlice(intsUsed)

	fmt.Print(intsUsed)
}
