package main

import "fmt"

func main() {
	//STRING
	// var greetingTexts string = "Hello World!"
	greetingTexts := "Hello World"
	fmt.Println(greetingTexts)
	fmt.Println(greetingTexts)

	//INT
	luckyNumber := 10
	anotherNumber := luckyNumber + 10
	fmt.Println(luckyNumber)
	fmt.Println(anotherNumber)
	anotherNumber = luckyNumber * 3
	fmt.Println(anotherNumber)

	//FLOAT 64
	var newNumber float64 = float64(luckyNumber) / 3
	fmt.Println(newNumber)

	//FLOAT 32
	var defaultFloat float64 = 1.165123554432316484873030
	var smallFloat float32 = 1.165123554432316484873030

	fmt.Println(defaultFloat)
	fmt.Println(smallFloat)

	//RUNE
	var character rune = '&'
	fmt.Println(string(character))

	//BYTE
	var firstByte byte = 'a'
	fmt.Println(firstByte) //97

	//FORMATING
	firstName := "Amar"
	lastName := "Nath"
	age := 23
	text := fmt.Sprintf("I am %v %v and my age is %v", firstName, lastName, age)
	//with the help of sprit we can store the output in some variable
	fmt.Println(text)

	const pi = 3.14
	//cannot reassign it

}
