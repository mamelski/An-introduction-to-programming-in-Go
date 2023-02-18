package main

import "fmt"

func average(xs []float64) (avg float64) {
	total := 0.0
	for _, v := range xs {
		total += v
	}

	avg = total / float64(len(xs))
	return
}

func multipleReturnValues() (a int, b int) {
	return 1, 2
}

func variadicFunction(args ...float64) float64 {
	total := 0.0
	for _, v := range args {
		total += v
	}
	return total
}

func main() {
	xs := []float64{98, 93, 77, 82, 83}

	fmt.Println(average(xs))
	fmt.Println(multipleReturnValues())

	fmt.Println(variadicFunction(xs...))
}
