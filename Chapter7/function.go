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

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func main() {
	xs := []float64{98, 93, 77, 82, 83}

	fmt.Println(average(xs))
	fmt.Println(multipleReturnValues())

	fmt.Println(variadicFunction(xs...))

	printer := func(number int) {
		fmt.Println("Printer", number)
	}

	x := 0
	increment := func() (result int) {
		x++
		return x
	}

	printer(increment())
	printer(increment())
	printer(increment())

	printer(1)
	printer(2)

	nextEven := makeEvenGenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())

	fmt.Println(factorial(6))
}
