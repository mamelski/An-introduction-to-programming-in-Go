package main

import "fmt"

func main() {
	x := [5]float64{13, 123, 123, 1, 14}

	var total float64 = 0
	for i, value := range x {
		fmt.Println(i, value)
		total += value
	}

	fmt.Println("total", total)

	average := total / float64(len(x))
	fmt.Println("average", average)

	fmt.Println(x[3])
}
