package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x float64
	y float64
	r float64
}

func circleArea(c Circle) float64 {
	return math.Pi * math.Pow(c.r, 2)
}

func main() {

	c := Circle{
		x: 0,
		y: 0,
		r: 5,
	}

	fmt.Println(c)
	fmt.Println(c.x)
	fmt.Println(c.y)
	fmt.Println(c.r)

	fmt.Println(circleArea(c))
}
