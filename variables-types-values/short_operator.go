package main

import (
	"fmt"
	"strconv"
)

var z = 30

func main() {
	age := 27
	message := "Hello, my name is Alejandro and I'm " + strconv.Itoa(age) + " years old."
	number := getNumber()

	fmt.Println(message)
	fmt.Println(number)
}

func getNumber() int {
	return z
}
