package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func main() {
	cow := Animal{food: "grass", locomotion: "walk", noise: "moo"}
	bird := Animal{food: "worms", locomotion: "fly", noise: "peep"}
	snake := Animal{food: "mice", locomotion: "slither", noise: "hsss"}

	animalMap := make(map[string]Animal)

	animalMap["cow"] = cow
	animalMap["bird"] = bird
	animalMap["snake"] = snake

	scanner := bufio.NewScanner(os.Stdin)
	for {
		var userInput string
		fmt.Print("> ")
		scanner.Scan()
		userInput = strings.ToLower(scanner.Text())

		inputSplice := strings.Split(userInput, " ")

		switch inputSplice[1] {
		case "eat":
			animalMap[inputSplice[0]].Eat()
		case "move":
			animalMap[inputSplice[0]].Move()
		case "speak":
			animalMap[inputSplice[0]].Speak()
		}
	}
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}
