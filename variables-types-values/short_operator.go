package main

import (
	"fmt"
	"strconv"
)

var z int = 30
var lastName string = "Hernández"

func main() {
	age := 27
	message := "Hello, my name is Alejandro and I'm " + strconv.Itoa(age) + " years old."
	number := getNumber()

	fmt.Println(message)
	fmt.Println(number)
	fmt.Println(lastName)
}

func getNumber() int {
	return z
}
