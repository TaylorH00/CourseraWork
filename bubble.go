package main

import "fmt"

func main() {
	ints := make([]int, 10)

	for i := range ints {
		fmt.Printf("Enter an integer (%d/10):", i+1)
		fmt.Scan(&ints[i])
	}

	BubbleSort(ints)

	fmt.Print(ints)
}

func BubbleSort(slice []int) {
	for range slice {
		swapped := false
		for i := 0; i < len(slice)-1; i++ {
			for j := 0; j < len(slice)-i-1; j++ {
				if slice[j] > slice[j+1] {
					Swap(slice, j)
				}
			}
		}
		if !swapped {
			break
		}
	}
}

func Swap(slice []int, index int) {
	currentVal := slice[index]
	slice[index] = slice[index+1]
	slice[index+1] = currentVal
}
