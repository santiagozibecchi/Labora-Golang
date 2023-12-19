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

type IntSequence interface {
	Next() int
	Title() string
}
type Predicate interface {
	Fulfill(n int) bool
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


// Implementa los metodos de la interface IntSequence ==> es decir, todos los structs de tipo Sequence pueden hacer uso de los metodos de la interface IntSequence
type Sequence struct {
	currentNumber 	int
	generator       string // Tipo de secuencia
	predicate       Predicate // evaluar tipo de secuencia 
	direction 		int // asc 0 - desc 1
}

func (s *Sequence) Next() int {
	var next int
	fulfillsPredicate := false
	var n int

	if s.direction == 0 {
		n = s.currentNumber + 1
	} else {
		n = s.currentNumber - 1
	}

	for !fulfillsPredicate {

		if s.predicate.Fulfill(n) {
			next = n
			fulfillsPredicate = true
		}

		if s.direction == 0 {
			n++
		} else {
			n--
		}
	}

	s.currentNumber = next
	return s.currentNumber
}

func (s *Sequence) Title() string {
	if s.direction == 0 {
		return "Secuencia de " + s.generator + " en Orden Ascendente:"
	} 

	return "Secuencia de " + s.generator + " en Orden Descendiente:"
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
	currentNumber := 30
	fmt.Println("El resultado de la tira del dado es: ", randomNumber)
	
	newSequense := []IntSequence{
		&Sequence{currentNumber: currentNumber, generator: "Números enteros", predicate: &LinearSequence{}, direction: zeroOrOne},
		&Sequence{currentNumber: currentNumber, generator: "Números primos", predicate: &PrimesSequence{}, direction: zeroOrOne},
		&Sequence{currentNumber: currentNumber, generator: "Números pares", predicate: &EvenSequence{}, direction: zeroOrOne},
	}

	selectedSequense := newSequense[randomNumber]
	
	fmt.Println(selectedSequense.Title())

	for i := 0; i < 30; i++ {
		fmt.Printf("%d ", selectedSequense.Next())
	}




}




