package main

import (
	"fmt"
	"math"
)

func  ejercicio01()  {
	/*
		1. Desarrolle una función para determinar si un número natural (> 0) es perfecto. 
		Un número es perfecto si la suma de los divisores propios da el el mismo numero. 
		* Los divisores propios de un número son todos sus divisores sin contar el mismo número.       
		* Sean X1, X2, XN todos divisores propios de X   
		* Si X es propio entonces X1 + x2 +…+ XN es igual a X
		
		Ejemplo:       
		**6 es un número perfecto**    
		**Divisores Propios: 1, 2 y 3.**
		**1 + 2 + 3 = 6**
	*/	
		enterNumber := 28
	
	if isPerfectNumber(enterNumber) {
		fmt.Printf("%d es un numero perfecto\n", enterNumber)
	} else {
		fmt.Printf("%d NO es un numero perfecto\n", enterNumber)
	}	
}

func isPerfectNumber(enterNumber int) bool {
	// modularizamos la lógica
	return posiblesDividers(enterNumber) == enterNumber
}

func posiblesDividers(enterNumber int) int {
	accDivider := 0
	// divisores posibles enterNumber/2
	for i := 1; i <= enterNumber/2; i++ {
		if enterNumber % i == 0 {
			accDivider += i
		}
	}
	return accDivider
}

func areFriendsNumber() bool {

	number1 := 220
	number2 := 284

	accDivider1 := posiblesDividers(number1)
	fmt.Println("accDivider1", accDivider1)
	accDivider2 := posiblesDividers(number2)
	fmt.Println("accDivider2", accDivider2)

	return accDivider2 == number1 && accDivider1 == number2
}


func ejercicio02()  {
	if areFriendsNumber() {
		fmt.Printf("Son numeros amigos\n")
	} else {
		fmt.Printf("NO Son numeros amigos\n")
	}
}

func esPrimo(enterNumber int ) bool {

	for i := 2; i < int(math.Sqrt(float64(enterNumber)))+1; i++ {

		if (enterNumber % i == 0) {
			return false
		} 
	}
	
	return true
}

func ejercicio03()  {
	
	if esPrimo(97)  {
		fmt.Println("Es primo")
	} else {
		fmt.Println("NO Es primo")
	}

}

func ejercicio08() string {
/*
	Realice un algoritmo que dado un string lo invierta.
*/
	str := "Hola muñdó!"
	runas := []rune(str) 
	var reverseStr []rune 

	fmt.Println(runas)

	for i := len(runas) - 1; i >= 0; i-- {
		reverseStr = append(reverseStr, runas[i])
	}
	return string(reverseStr)
}

func main() {

	// ejercicio01()
	// ejercicio02()
	// ejercicio03()


	result := ejercicio08()
	fmt.Println(result)

}