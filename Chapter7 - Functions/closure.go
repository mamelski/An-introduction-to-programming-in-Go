package main

import "fmt"

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func main() {

	x := 0
	increment := func() (result int) {
		x++
		return x
	}

	fmt.Println(increment())
	fmt.Println(increment())

	fmt.Println(increment())

	nextEven := makeEvenGenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())

}
