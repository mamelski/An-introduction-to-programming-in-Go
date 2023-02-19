package main

import "fmt"

func sum(numbers ...int) int {

	result := 0
	for _, v := range numbers {
		result += v
	}

	return result
}

func main() {
	numbers := []int{12, 342, 345, 4, 5, 6, 7, 8}

	sum := sum(numbers...)

	fmt.Println(sum)
}
