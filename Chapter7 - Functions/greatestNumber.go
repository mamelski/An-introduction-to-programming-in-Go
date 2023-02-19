package main

import (
	"fmt"
	"math"
)

func greatestNumber(numbers ...float64) float64 {
	greatest := 0.0
	for _, v := range numbers {
		greatest = math.Max(greatest, v)
	}
	return greatest
}

func main() {

	numbers := []float64{1, 2, 34234, 4536, 8, 7.56, 5, 6, 7}

	fmt.Println(greatestNumber(numbers...))
}
