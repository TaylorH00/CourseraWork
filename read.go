package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Names struct {
	fname string
	lname string
}

func main() {
	var fileName string
	namesSlice := make([]Names, 0)

	fmt.Print("Enter the filename: ")
	fmt.Scan(&fileName)

	file := make([]byte, 1)

	file, _ = ioutil.ReadFile(fileName)

	tempSlice := strings.Split(string(file), "\n")

	for _, v := range tempSlice {
		if v != "" {
			temperSlice := strings.Split(v, " ")
			namesSlice = append(namesSlice, Names{fname: temperSlice[0], lname: temperSlice[1]})
		}
	}

	for _, name := range namesSlice {
		fmt.Printf("First Name: %s Last Name: %s \n", name.fname, name.lname)
	}
}
