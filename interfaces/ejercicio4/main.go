package main

/*
Ejercicio 4:

¿Qué pasaría si quiero una secuencia pero en el orden inverso? 🤔. 
¿Acaso es mucho pedir…? Creo que estoy teniendo sed de combinar dos capacidades: 
por un lado el orden en el que se generan los números: creciente o decreciente y, por otro lado, 
un predicado que me diga si un número cumple con tal condición.

Se desea implementar una estructura Sequence que tenga como campos un generador
y un predicado y que tenga un método llamado Next() que devuelva el próximo número
que se genera usando el generador y que cumpla con el predicado.
*/


import (
	"fmt"
	"math/rand"
	"time"
)

func isEven(n int) bool {
	if n % 2 == 0 {
		return true
	}
	return false
}

func isPrime(n int) bool {
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

type LinearSequence struct {}
type PrimesSequence struct {}
type EvenSequence struct {}

func (s* LinearSequence) Fulfill(n int) bool {
	return true
}
func (s* PrimesSequence) Fulfill(n int) bool {
	return isPrime(n)
}
func (s* EvenSequence) Fulfill(n int) bool {
	return isEven(n)
}

type Predicate interface {
	Fulfill(n int) bool
}

type Sequence struct {
	currentNumber 	int
	generator       string // Tipo de secuencia
	predicate       Predicate // evaluar tipo de secuencia 
	direction 		int // asc 0 - desc 1
}


func (s *Sequence) Next() int {
	s.currentNumber++

	

	return s.currentNumber
}

func (s *Sequence) Title() string {
	if s.direction == 0 {
		return "Secuencia" + s.generator + " Ascendente."
	} 
	
	return "Secuencia" + s.generator +" Descendiente."
}



func determineOrder(n int) bool {
	if n == 0 {
		return true 
	}

	return false
}


func main() {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(3)
	zeroOrOne := rand.Intn(2)
	var ascOrDescOrder bool
	fmt.Println("El resultado de la tira del dado es: ", randomNumber)

	ascOrDescOrder = determineOrder(zeroOrOne)
	if ascOrDescOrder {
		fmt.Println("Se armara una secuencia en orden: ASCENDENTE")
		fmt.Println("--------------------------------------------")
	} else {
		fmt.Println("Se armara una secuencia en orden: DESENDENTE")
		fmt.Println("--------------------------------------------")
	}

	





	for i := 0; i < 30; i++ {
		// fmt.Printf("%d ", sequence.Next())
	}
}
