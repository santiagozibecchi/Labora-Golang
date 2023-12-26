package main

import (
	"encoding/json"
	"fmt"
)

type persona struct {
	Nombre string
	Apellido string
	Edad int
}

func main() {
	
	p1 := persona{
		Nombre: "Santiago",
		Apellido: "Zibecchi",
		Edad: 28,
	}

	p2 := persona{
		Nombre: "Maria",
		Apellido: "Zibecchi",
		Edad: 28,
	}

	personas := []persona{p1, p2}
	fmt.Println(personas)

	// devuelve un slice de bits
	bs, err := json.Marshal(&personas)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println(string(bs))

}