package main

import (
	"fmt"
)

var h = "hello"

func main() {
	var x string
	x = "first"
	fmt.Println(x)
	x += "second"
	fmt.Println(x)

	var y = "hello"
	var z = "world"
	fmt.Println(y == z)

	z = "hello"
	fmt.Println(y == z)

	e := 5
	fmt.Println(e)

	dogsName := "Max"
	fmt.Println(dogsName)

	fmt.Println(h)

	const cannotChangeIt = "const"
	fmt.Println(cannotChangeIt)

	var (
		a = 7
		b = 8
		c = 9
	)

	fmt.Println(a, b, c)

	fmt.Println("Enter a number")
	var input float64
	fmt.Scanf("%f", &input)
	output := input * 2
	fmt.Println(output)
}
