package main

import "fmt"

func multipleReturnValues() (a int, b int) {
	return 1, 2
}

func main() {
	fmt.Println(multipleReturnValues())
}
