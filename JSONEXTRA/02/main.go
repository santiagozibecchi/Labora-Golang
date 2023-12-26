package main

import (
	"encoding/json"
	"fmt"
)

type persona struct {
	Nombre   string `json:"Nombre"`
	Apellido string `json:"Apellido"`
	Edad     int    `json:"Edad"`
}

func main() {
	
	s := `[{"Nombre":"Santiago","Apellido":"Zibecchi","Edad":28},{"Nombre":"Maria","Apellido":"Zibecchi","Edad":28}]`

	bs := []byte(s)

	fmt.Printf("%T\n", s)
	// byte => es una alias para uint8
	fmt.Printf("%T\n", bs)

	var personas []persona

	err := json.Unmarshal(bs, &personas)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("Toda la data: ", personas)

	for i, v := range personas {
		fmt.Println("\nPersona N°", i+1)
		fmt.Println(v.Nombre, v.Apellido, v.Edad )
	}
}