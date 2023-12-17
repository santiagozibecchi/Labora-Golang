package main

import "fmt"

type Figura interface {
	Area() float64
	Perimetro() float64
}

type Cuadrado struct {
	lado float64
}

type Circulo struct {
	radio float64
}

// Se implementa de forma implicita los metodos de la interfase Figura sobre los structs
func (c *Cuadrado) Area() float64 {
	return c.lado * c.lado
}
func (c *Cuadrado) Perimetro() float64 {
	return c.lado * 4
}

func (c *Circulo) Area() float64 {
	return 3.14*c.radio*c.radio
}
func (c *Circulo) Perimetro() float64 {
	return 2*3.14*c.radio
}

func main()  {
	
	cuadrado := Cuadrado{16}
	fmt.Println("Cuadrado", cuadrado)
	fmt.Println(cuadrado.Area())
	fmt.Println(cuadrado.Perimetro())

	circulo := Circulo{5}
	fmt.Println("Circulo", circulo)
	fmt.Println(circulo.Area())
	fmt.Println(circulo.Perimetro())


	fmt.Println("Area Ciculo: ", Figura.Area(&circulo))
	fmt.Println("Area Cuadrado: ", Figura.Area(&cuadrado))
}