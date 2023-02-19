package main

import "fmt"

func half(number int) (half int, isEven bool) {
	half = number / 2
	isEven = number%2 == 0
	return
}

func main() {
	fmt.Println(half(1))
	fmt.Println(half(2))
	fmt.Println(half(3))
	fmt.Println(half(4))
}
