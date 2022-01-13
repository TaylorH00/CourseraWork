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

func (c Cow) Eat() {
	fmt.Print("grass")
}
func (c Cow) Move() {
	fmt.Print("walk")
}
func (c Cow) Speak() {
	fmt.Print("moo")
}

type Bird struct{}

func (b Bird) Eat() {
	fmt.Print("worms")
}
func (b Bird) Move() {
	fmt.Print("fly")
}
func (b Bird) Speak() {
	fmt.Print("peep")
}

type Snake struct{}

func (s Snake) Eat() {
	fmt.Print("mice")
}
func (s Snake) Move() {
	fmt.Print("slither")
}
func (s Snake) Speak() {
	fmt.Print("hsss")
}

func AnimalCommand(a Animal, action string) {
	cow, ok := a.(Cow)
	if ok {
		switch action {
		case "eat":
			cow.Eat()
		case "move":
			cow.Move()
		case "speak":
			cow.Speak()
		}
	}
	bird, ok := a.(Bird)
	if ok {
		switch action {
		case "eat":
			bird.Eat()
		case "move":
			bird.Move()
		case "speak":
			bird.Speak()
		}
	}
	snake, ok := a.(Snake)
	if ok {
		switch action {
		case "eat":
			snake.Eat()
		case "move":
			snake.Move()
		case "speak":
			snake.Speak()
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	animalMap := make(map[string]interface{})

	for {
		var userInput string
		fmt.Print("> ")
		scanner.Scan()
		userInput = strings.ToLower(scanner.Text())

		inputSplice := strings.Split(userInput, " ")

		if inputSplice[0] == "newanimal" {
			switch inputSplice[3] {
			case "eat":
				animalMap[inputSplice[1]] = type Cow
			case "move":
				animalMap[inputSplice[1]] = Bird
			case "snake":
				animalMap[inputSplice[1]] = Snake
			}
			fmt.Print("Created It!")
		} else if inputSplice[0] == "query" {
			if animalMap[inputSplice[1]] != nil {
				AnimalCommand(animalMap[inputSplice[1]], inputSplice[2])
			} else {
				fmt.Print("No animal with that name exists.")
			}
		} else {
			fmt.Print("Please enter a valud command (\"newanimal\" or \"query\")")
		}
	}
}
