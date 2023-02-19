package main

import "fmt"

func makeOddNumbersGenerator() func() int {
	oddNumber := -1
	return func() int {
		oddNumber += 2
		return oddNumber
	}
}

func main() {
	nextOddNumber := makeOddNumbersGenerator()

	fmt.Println(nextOddNumber())
	fmt.Println(nextOddNumber())
	fmt.Println(nextOddNumber())
	fmt.Println(nextOddNumber())
	fmt.Println(nextOddNumber())
	fmt.Println(nextOddNumber())

}
