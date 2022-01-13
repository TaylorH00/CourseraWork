package main

import (
	"fmt"
	"time"
)

var sharedInt int = 0

func Function1() {
	for i := 0; i < 3; i++ {
		fmt.Println("Printer: ", sharedInt)
	}
}

func Function2() {
	for i := 0; i < 3; i++ {
		sharedInt = sharedInt + 1
		fmt.Println("Setter: ", sharedInt)
	}
}

func main() {
	go Function1()
	go Function2()
	time.Sleep(time.Second)
	fmt.Println("Done")
}
