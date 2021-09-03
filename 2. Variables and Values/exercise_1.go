package variable

import "fmt"

func main() {
	var firstName string = "Amar"
	lastName := "Nath"
	fmt.Println(firstName, lastName)
	fmt.Println(lastName)
	var currentYear int = 2021
	birthYear := 1999
	fmt.Println(currentYear - birthYear)
	currentYear = currentYear + 1
	fmt.Println(currentYear - birthYear)
}
