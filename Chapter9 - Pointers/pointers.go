package main

import "fmt"

func zero(x int) {
	x = 0
}

func zeroPtr(xPtr *int) {
	*xPtr = 0
}

func main() {
	x := 5
	zero(x)
	fmt.Println(x)

	zeroPtr(&x)
	fmt.Println(x)

}
