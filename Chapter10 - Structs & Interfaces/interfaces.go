package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func distance3(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

type Circle3 struct {
	x float64
	y float64
	r float64
}

type Rectangle3 struct {
	x1 float64
	x2 float64
	y1 float64
	y2 float64
}

func (r *Rectangle3) area() float64 {
	l := distance3(r.x1, r.y1, r.x2, r.y2)
	w := distance3(r.x1, r.y1, r.x2, r.y2)
	return l * w
}

func (c *Circle3) area() float64 {
	return math.Pi * math.Pow(c.r, 2)
}

type Multishape struct {
	shapes []Shape
}

func (m *Multishape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}

func main() {

	c := Circle3{
		x: 1,
		y: 2,
		r: 4,
	}

	r := Rectangle3{
		x1: 1,
		x2: 2,
		y1: 3,
		y2: 4,
	}

	fmt.Println(totalArea(&c, &r))
}
