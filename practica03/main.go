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
	return posiblesOwnDividers(enterNumber) == enterNumber
}

func posiblesOwnDividers(enterNumber int) int {
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

	accDivider1 := posiblesOwnDividers(number1)
	fmt.Println("accDivider1", accDivider1)
	accDivider2 := posiblesOwnDividers(number2)
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

// func isPrime(n int) {
// 	numerosPropios(n) == 1 
// }

func ejercicio03()  {
	
	if esPrimo(97)  {
		fmt.Println("Es primo")
	} else {
		fmt.Println("NO Es primo")
	}

}

func ejercicio07() string {
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

func  ejercicio04()  {
	/*
		El algoritmo más eficiente es prime_v2, ya que una vez que se cumple la condición de NO ser un numero primo
		rompe la condición del ciclo for y no sigue iterando hasta n como lo hace el algoritmo prime_v1
	*/
}

func ejercicio05 () {
	var cadena string
	fmt.Print("Ingrese un número como cadena: ")
	fmt.Scan(&cadena)

	resultado := toInt(cadena)
	fmt.Printf("El resultado como entero es: %d\n", resultado)
}

func toInt(cad string) int {
	// longuitud de la cadena
	long := len(cad)
	i, pot, res := 0, 1, 0

/*
	Para empezar:
	ref: https://zetcode.com/golang/byte/
	A byte in Go is an unsigned 8-bit integer. It has type uint8. 
	A byte has a limit of 0 – 255 in numerical range => It can represent an ASCII character.

	El valor numérico del carácter '0' en esta tabla ASCII es 48 (decimal).
	48  49  50  51  52  53  54  55  56  57
	'0' '1' '2' '3' '4' '5' '6' '7' '8' '9'

*/

	for i < long {
/* 
	La idea principal se trata de sumar numeros del 1 al 9.
	Utilizamos los valores ASCII para convertir cada carácter a su valor numérico.
	
	Se obtiene cada byte en orden descendente (long-i-1) para sumar del último al primero, 
	teniendo en cuenta la cantidad de ceros basada en su posición mediante la potencia (pot).
*/
		c := cad[long-i-1]
		res += int(c-'0') * pot
		pot *= 10
		i++
	}

	return res
}

func ejercicio06()  {
	// Comenta que realiza el siguiente algoritmo. Una vez entendido lo que hace el algoritmo 
	// proponga nuevos nombres para los identificadores que describan mejor su uso.
	// https://go.dev/play/p/3-c0WgRrXDO
}

func ejercicio08()  {
	/*
	Realice un algoritmo que dado un string te diga si es palindromo
	*/

}

func ejercicio09()  {
	/*
	Realice un algoritmo que dado un número te diga si es capicua.
	*/
	// RESUELTO ANTERIORMENTE
}

func ejercicio10()  {
	/*
	Realice un algoritmo para multiplicar el factorial de un número por su sumatoria.
	*/
}

func main() {

	// ejercicio01()
	// ejercicio02()
	// ejercicio03()

	// ejercicio04()
	// ejercicio05()
	// ejercicio06()
	
	
	// result := ejercicio07()
	// fmt.Println(result)
	
	// TODO 
	// ejercicio08()
	// Resuelto pero no acá ejercicio09()
	// ejercicio10()


}