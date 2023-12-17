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
*/

type IntSequence interface {
	Next() int
}

// structs encargados de manejar el conteo condicional
type IncreasingLinearSequence struct {
	currentNumber int
}

type IncreasingPrimeSequence struct {
	currentNumber int
}

func(s* IncreasingLinearSequence) Next() int {
	s.currentNumber++
	return s.currentNumber
}

func (s* IncreasingPrimeSequence) Next() int {

	s.currentNumber++

	for !isPrime(s.currentNumber) {
		s.currentNumber++
	}

	return s.currentNumber
}

func isPrime(n int) bool {
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}



func main()  {
	
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(2)
	fmt.Println("El resultado de la tira del dado es: ", randomNumber)

	intSequence := []IntSequence{&IncreasingLinearSequence{}, &IncreasingPrimeSequence{}}

	for i := 0; i < 30; i++ {
		fmt.Printf("%d ", intSequence[randomNumber].Next())
	}

}