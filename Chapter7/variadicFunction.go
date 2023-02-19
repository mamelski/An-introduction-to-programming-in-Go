package main

import "fmt"

func variadicFunction(args ...float64) float64 {
	total := 0.0
	for _, v := range args {
		total += v
	}
	return total
}

func main() {
	xs := []float64{98, 93, 77, 82, 83}

	fmt.Println(variadicFunction(xs...))

}
