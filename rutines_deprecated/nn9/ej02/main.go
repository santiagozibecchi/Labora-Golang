package main

import "fmt"

type humano interface {
	hablar()
}

type persona struct {
	nombre string
}

func (p *persona) hablar() {
	fmt.Println("Hola! Soy ", p.nombre)
}

func diAlgo(h humano) {
	h.hablar()
}

func main()  {
	
	p1 := persona{
		nombre: "Santiago",
	}

	diAlgo(&p1)

	p1.hablar()

}