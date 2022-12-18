package main

import "fmt"

func main() {

	fmt.Println("Enter Fahrenheit degrees:")
	var fahrenheit float64
	fmt.Scanf("%f", &fahrenheit)

	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Println(celsius)

}
