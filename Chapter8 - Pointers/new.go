package main

import "fmt"

func one(xPtr *int) {
	*xPtr = 1
}

func square(x *float64) {
	*x = *x * *x
}

func main() {
	// new creates a pointer
	pointer := new(int)

	one(pointer)
	fmt.Println(*pointer)

	x := 1.5
	square(&x)
	fmt.Println(x)
}
