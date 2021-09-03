package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	fmt.Println("BMI Calculator")
	fmt.Println("----------------")

	fmt.Print("Enter the weight of the person (in Kg): ")
	weightInput, _ := reader.ReadString('\n')

	fmt.Print("Enter the height of the person (in m): ")
	heightInput, _ := reader.ReadString('\n')

	fmt.Print(weightInput)
	fmt.Print(heightInput)
	
	//for windows \r\n
	weightInput = strings.Replace(weightInput, "\r\n", "", -1)
	heightInput = strings.Replace(heightInput, "\r\n", "", -1)

	weight, _ := strconv.ParseFloat(weightInput, 64)
	height, _ := strconv.ParseFloat(heightInput, 64)

	bmi := weight / (height * height)

	fmt.Printf("The BMI is %.2f", bmi)
}
