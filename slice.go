package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	intSlice := make([]int, 3)

	var userInput string

	for userInput != "X" {
		fmt.Print("Enter an Integer:        (Enter X to exit.)")
		fmt.Scan(&userInput)

		i, _ := strconv.Atoi(userInput)

		intSlice = append(intSlice, i)

		sort.Ints(intSlice)

		fmt.Print(intSlice)
	}
}
