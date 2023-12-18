package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Ejercicio 1:

Se necesita generar dos secuencias de números que sigan la siguiente lógica
1. Números en orden crecientes (del 1 en adelante)
2. Números que sean primos (1,2,3,5,7,9,11…)
Se pide implementar cada secuencia en una struct que siga cumpla con la interfaz:
type IntSequence interface {

Next() int

}

Por último se pide desarrollar un pequeño programa donde se pueda imprimir
los primeros 30 números de una de las dos secuencias según una “tirada de dados”.

Tip: Podemos simular una tirada de dado usando rand.Intn(2) (del paquete math/rand).

Ejercicio 2:
Ahora se desea que se imprima un mensaje que diga cuál de las dos secuencias se tomó (según la tirada de dados).

TIP Se le puede declarar un método Title() string en la interface IntSequence y hacer que ambas secuencias lo implementen.

Ejercicio 3:

De forma más general se desea tener una secuencia que me dé números que cumplan una condición (predicado).
Dicha condición será implementada en estructuras (o cualquier otro tipo de dato)  que tengan el un método con la siguiente firma `Fulfill(n int) bool`.
 ***Fulfill*** es un término en inglés que comúnmente se emplea para afirmar que algo cumple una condición.
Además de tener tipos de datos que tengan una función para implementar condiciones de:
Números cualquier sea (para secuencias de enteros lineales)
Números primos
Ahora quiero que agreguen un nuevo tipo que tenga una función que implemente la condición de
Números pares (2,4,6…)
*/


type IntSequence interface {
	Next() int
	Title() string
	Fulfill(n int) bool
}

// structs encargados de manejar el conteo condicional
type IncreasingLinearSequence struct {
	currentNumber int
}

func(s* IncreasingLinearSequence) Next() int {	

	if s.Fulfill(s.currentNumber) {
		s.currentNumber++
	}

	return s.currentNumber
}

func (s* IncreasingLinearSequence) Title() string {
	return "Secuencia Lineal de Numeros Enteros: "
}

func (s* IncreasingLinearSequence) Fulfill(n int) bool {
	return true
}

type IncreasingPrimeSequence struct {
	currentNumber int
}

func (s* IncreasingPrimeSequence) Next() int {

	s.currentNumber++

	for !s.Fulfill(s.currentNumber) {
		s.currentNumber++
	}

	return s.currentNumber
}

func (s* IncreasingPrimeSequence) Title() string {
	return "Secuencia de Numeros Primos: "
}

func (s* IncreasingPrimeSequence) Fulfill(n int) bool {
	return isPrime(n)
}

func isPrime(n int) bool {
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Secuencia de Numeros pares:
type IncreasingEvenSequence struct {
	currentNumber int
}

func (s* IncreasingEvenSequence) Title() string {
	return "Secuencia de Numeros Pares: "
}

func (s* IncreasingEvenSequence) Fulfill(n int) bool {
	return isEven(n)
}

func (s* IncreasingEvenSequence) Next() int {
	s.currentNumber++

	for !s.Fulfill(s.currentNumber){
		s.currentNumber++
	}

	return s.currentNumber
}

func isEven(n int) bool {
	if n % 2 == 0 {
		return true
	}
	return false
}

func main()  {
	
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(3)
	fmt.Println("El resultado de la tira del dado es: ", randomNumber)

	intSequence := []IntSequence{&IncreasingLinearSequence{}, &IncreasingPrimeSequence{}, &IncreasingEvenSequence{}}

	selectedSequence := intSequence[randomNumber]
	fmt.Println(selectedSequence.Title())

	for i := 0; i < 30; i++ {
		fmt.Printf("%d ", intSequence[randomNumber].Next())
	}

}