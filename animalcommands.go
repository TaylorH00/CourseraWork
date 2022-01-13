package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{}
type Bird struct{}
type Snake struct{}

func (c Cow) Eat() {
	fmt.Println("grass")
}
func (c Cow) Move() {
	fmt.Println("walk")
}
func (c Cow) Speak() {
	fmt.Println("moo")
}

func (b Bird) Eat() {
	fmt.Println("worms")
}
func (b Bird) Move() {
	fmt.Println("fly")
}
func (b Bird) Speak() {
	fmt.Println("peep")
}

func (s Snake) Eat() {
	fmt.Println("mice")
}
func (s Snake) Move() {
	fmt.Println("slither")
}
func (s Snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	animalMap := make(map[string]Animal)
	scan := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(">")
		scan.Scan()
		userInput := strings.Split(scan.Text(), " ")

		if userInput[0] == "newanimal" {
			if !(strings.Contains(userInput[2], "cow") ||
				strings.Contains(userInput[2], "bird") ||
				strings.Contains(userInput[2], "snake")) {
				fmt.Println("Enter a valid animal type.")
				continue
			}
			switch userInput[2] {
			case "cow":
				animalMap[userInput[1]] = Cow{}
			case "bird":
				animalMap[userInput[1]] = Bird{}
			case "snake":
				animalMap[userInput[1]] = Snake{}
			}
			fmt.Println("Created it!")
		} else if userInput[0] == "query" {
			if !(strings.Contains(userInput[2], "eat") ||
				strings.Contains(userInput[2], "move") ||
				strings.Contains(userInput[2], "speak")) {
				fmt.Println("Enter a valid information type.")
				continue
			}
			animal := animalMap[userInput[1]]
			switch userInput[2] {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			}
		} else {
			fmt.Println("Enter a valid command.")
		}

	}
}
