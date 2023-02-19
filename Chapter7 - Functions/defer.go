package main

import "fmt"

func first() {
	fmt.Println("first")
}

func second() {
	fmt.Println("second")
}

func main() {
	// defer is often use to free resources (like closing file at the end of function)
	// defer runs even if run-time panic occurs
	defer second()
	first()
}
