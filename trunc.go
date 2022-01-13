package main

import (
	"fmt"
	"math"
)

func main() {
	var intValue int

	fmt.Print("Enter a floating point number:")

	var floatingValue float32

	fmt.Scan(&floatingValue)

	intValue = int(math.Trunc(float64(floatingValue)))

	fmt.Println("Here is the truncated version of your floating point number:", intValue)
}
