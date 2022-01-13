package main

import (
	"fmt"
	"math"
)

func GenDisplaceFn(accel, initVel, initDisp float64) func(t float64) float64 {
	fn := func(t float64) float64 {
		return (0.5*accel)*(math.Pow(t, 2)) + (initVel * t) + initDisp
	}

	return fn
}

func main() {
	var acceleration float64
	var initVelocity float64
	var initDisplacement float64
	var time float64

	fmt.Print("Enter acceleration:")
	fmt.Scan(&acceleration)
	fmt.Print("Enter initial velocity:")
	fmt.Scan(&initVelocity)
	fmt.Print("Enter initial displacement:")
	fmt.Scan(&initDisplacement)
	fmt.Print("Enter time:")
	fmt.Scan(&time)

	fn := GenDisplaceFn(acceleration, initVelocity, initDisplacement)

	fmt.Println(fn(time))
}
