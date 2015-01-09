package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %g", e)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	startingPoint := x
	change := 1.0
	for change > 1e-14 {
		newStartingPoint := startingPoint - (((startingPoint * startingPoint) - x) / (2 * startingPoint))
		change = math.Abs(newStartingPoint - startingPoint)
		startingPoint = newStartingPoint
	}
	return startingPoint, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
