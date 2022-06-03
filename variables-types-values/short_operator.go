package main

import (
	"fmt"
	"strconv"
)

func main() {
	age := 27
	message := "Hello, my name is Alejandro and I'm " + strconv.Itoa(age) + " years old."

	fmt.Println(message)
}
