package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	startingPoint := x
	change := 1.0
	for change > 1e-16 {
		newStartingPoint := startingPoint - (((startingPoint * startingPoint) - x) / (2 * startingPoint))
		change = math.Abs(newStartingPoint - startingPoint)
		startingPoint = newStartingPoint
	}
	return startingPoint
}

func main() {
	x := float64(1337)
	fmt.Println(Sqrt(x), "versus", math.Sqrt(x))
}
