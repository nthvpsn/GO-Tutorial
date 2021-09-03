package main

import (
	"fmt"
	"math/rand"
)

func main() {
	
	sum := sum(5, 12)
	fmt.Println(sum)
	
	printNumber(5)
	
	r1, r2 := returnMulti()
	println(r1, r2)
	
	fmt.Println(nameReturn(5, 10))
}

//funtion to return the sum of two numbers
func sum(num1 int, num2 int) int {
	return num1 + num2
}

//func with no return
func printNumber(nums int) {
	fmt.Println(nums)
}

//multiple return value
func returnMulti() (int, int) {
	r1 := rand.Intn(10)
	r2 := rand.Intn(10)
	return r1, r2
}

//naming the return value
func nameReturn(num1 int, num2 int) (rt1 int, rt2 int) {
	rt1 = num1 * num1
	rt2 = num2 * 5
	return
}
