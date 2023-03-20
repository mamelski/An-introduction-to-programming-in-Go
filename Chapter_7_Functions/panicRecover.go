package main

import "fmt"

func main() {

	defer func() {
		error := recover()
		fmt.Println(error)
	}()

	panic("PANIC")
}
