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

	fmt.Println(personas)

	bs, err := json.Marshal(personas)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

	err = json.Unmarshal(bs, &personas)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Toda la data", personas)

	for i, v := range personas {
		fmt.Println("\nPERSONA NÚMERO", i+1)
		fmt.Println(v.Nombre, v.Apellido, v.Edad)
	}
}
