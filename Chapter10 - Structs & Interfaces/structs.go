package main

import "fmt"

type Circle struct {
	x float64
	y float64
	r float64
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
}
