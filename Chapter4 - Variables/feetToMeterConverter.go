package main

import "fmt"

func main() {
	fmt.Println("Enter distance in feet:")
	var distanceInFeet float64
	fmt.Scanf("%f", &distanceInFeet)

	distanceInMeters := distanceInFeet * 0.3048
	fmt.Println(distanceInMeters)
}
