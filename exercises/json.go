package main

import (
	"encoding/json"
	"fmt"
)

type persona struct {
	Nombre string
	Apellido  string
	Edad   int
}

func main() {
	p1 := persona{
		Nombre: "James",
		Apellido:  "Bond",
		Edad:   32,
	}

	p2 := persona{
		Nombre: "Miss",
		Apellido:  "Moneypenny",
		Edad:   27,
	}

	var personas []persona = []persona{p1, p2}

	fmt.Println(personas)// => Return JSON value

	bs, err := json.Marshal(personas) // => Convert to Byte String
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))// => Return Byte String

	err = json.Unmarshal(bs, &personas)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Toda la data", personas)// => Return JSON value

	for i, v := range personas {
		fmt.Println("\nPERSONA NÚMERO", i+1)
		fmt.Println(v.Nombre, v.Apellido, v.Edad)
	}
}
