package main

import "fmt"

func swap(xPtr *int, yPtr *int) {
	tmp := *xPtr
	*xPtr = *yPtr
	*yPtr = tmp
}

func main() {

	x, y := 1, 2
	swap(&x, &y)

	fmt.Println(x, y)
}
