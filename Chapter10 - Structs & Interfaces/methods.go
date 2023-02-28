package main

import (
	"fmt"
	"math"
)

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

type Circle2 struct {
	x float64
	y float64
	r float64
}

type Rectangle struct {
	x1 float64
	x2 float64
	y1 float64
	y2 float64
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x2, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y2)
	return l * w
}

func (c *Circle2) area() float64 {
	return math.Pi * math.Pow(c.r, 2)
}

func main() {

	c := Circle2{
		x: 0,
		y: 0,
		r: 5,
	}

	fmt.Println(c.area())

	r := Rectangle{1, 3, 10, 10}
	fmt.Println(r.area())
}
