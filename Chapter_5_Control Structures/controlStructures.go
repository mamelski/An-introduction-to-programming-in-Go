package main

import "fmt"

func main() {

	fmt.Println(`
		1
		2`)

	i := 1
	for i <= 10 {
		if i%2 == 0 {
			println(i, "even")
		} else {
			println(i, "odd")
		}
		i += 1
	}

	for j := 11; j <= 20; j++ {
		if j%2 == 0 {
			println(i, "even")
		} else {
			println(i, "odd")
		}
	}

	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	}
}
