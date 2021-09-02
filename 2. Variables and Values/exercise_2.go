package main

import "fmt"

func main() {
	PI := 3.14
	radius := 5

	cirumference := 2 * PI * float64(radius)
	fmt.Println(cirumference)

	fmt.Printf("For a radius of %v, the circle circumference is %.2f", radius, (cirumference))
}
