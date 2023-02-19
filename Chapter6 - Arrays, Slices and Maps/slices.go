package main

import (
	"fmt"
	"reflect"
)

func main() {

	x := make([]float64, 5)
	fmt.Println(x, reflect.TypeOf(x))

	y := make([]float64, 5, 10)
	fmt.Println(y, reflect.TypeOf(y), len(y), cap(y))

	arr := []float64{1, 2, 3, 4, 5}
	z := arr[0:2]
	fmt.Println(z)

	a := arr[len(arr)-2:]
	fmt.Println(a)

	a1 := append(z, 3)
	a2 := append(a1, a...)
	fmt.Println(a2)

	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 3, 4)
	copy(slice2, slice1)
	fmt.Println(slice2)

	slice3 := make([]int, 3, 9)
	fmt.Println(cap(slice3))
	fmt.Println(len(slice3))

	x2 := [6]string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(x2[2:5])

}
