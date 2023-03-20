package main

import "fmt"
import m "golang-book/Chapter_11_Packages/math2"

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := m.Average(xs)

	fmt.Println(avg)

	fmt.Println(m.Min(xs))
	fmt.Println(m.Max(xs))
}
