package main

import (
	"fmt"
	"sort"
)

type Persona struct {
	Nombre string
	Edad   int
}

type PorEdad []Persona

func (a PorEdad) Len() int            { return len(a) }
func (a PorEdad) Swap(i, j int)     { a[i], a[j] = a[j], a[i] }
func (a PorEdad) Less(i, j int) bool { return a[i].Edad < a[j].Edad }

func main() {
	p1 := Persona{"James", 32}
	p2 := Persona{"Moneypenny", 27}
	p3 := Persona{"Q", 64}
	p4 := Persona{"M", 56}

	personas := []Persona{p1, p2, p3, p4}

	fmt.Println(personas)
	sort.Sort(PorEdad(personas))
	fmt.Println(personas)
}
