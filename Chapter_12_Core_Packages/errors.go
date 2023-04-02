package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("nie działa")
	fmt.Println(err)
}
