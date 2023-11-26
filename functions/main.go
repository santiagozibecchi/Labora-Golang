package main

import (
	"fmt"
	"math"
	"strconv"
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

	enterString := "neuquen"

	if reverseString(enterString) == enterString {
		fmt.Println("Es string ingresado SI es palindromo")
	} else {
		fmt.Println("Es string ingresado NO es palindromo")
	}

}

func reverseString(input string) string {

    runes := []rune(input)
    length := len(runes)

    for i := 0; i < length/2; i++ {
        // Intercambiar caracteres desde el principio con los del final
        runes[i], runes[(length-1)-i] = runes[(length-1)-i], runes[i]
    }

    return string(runes)
}

func isCapicua(enterfield string) bool {

	if reverseString(enterfield) == enterfield {
		return true
	}

	return false
}

func ejercicio09(){
 /*
	Dado un número de 5 cifras, determinar si es capicúa. Si fuera un número de 6 cifras 
	¿Sirve la resolución planteada? ¿Cómo habría que modificarla?
 */

	enterNumber := 12321
	stringCovertedNumber := strconv.Itoa(enterNumber)


	if isCapicua(stringCovertedNumber) {
		fmt.Printf("El numero %d es capicua\n", enterNumber)
	} else {
		fmt.Printf("El numero %d, no es capicua\n", enterNumber)
	}

}

func ejercicio10()  {
	/*
	Realice un algoritmo para multiplicar el factorial de un número por su sumatoria.
	*/
	/*
	Realice un algoritmo para multiplicar el factorial de un número por su sumatoria.
*/
	enterNumber := 3

	multiplicationFacXSum := sumatoria(enterNumber) * factorial(enterNumber)

	fmt.Println("factorial de un número por su sumatoria", multiplicationFacXSum)
}

func factorial(number int) int {

	if number == 0 {
		return 1		
	}
	
	return number * factorial(number-1)
}

func sumatoria(number int) int {

	acc := 0
	for i := number; i > 0; i-- {
		acc+=i
	}

	return acc
}


func potencia(x int, y int) int  {
	acc := 1

	for i := 0; i < y; i++ {
		acc*=x
	}

	return acc
}

func mcm(number1 int, number2 int) int {
	maxCount := number1*number2
	mcm := 0

	for i := 1; i <= maxCount; i++ {
		
		if number1%i == 0 && number2%i == 0 {
			mcm = i
			break
		}
	}
	return mcm
}

func ejercicio11()  {
	/*
		Realice un algoritmo que dado dos números calcule el resultado de la potencia del primero elevado 
		al segundo más la sumatoria del primero multiplicado el segundo, todo lo anterior dividido
		el mínimo común múltiplo entre ambos números.
	*/

	var x int
	var y int
	var result int

	fmt.Print("Ingrese el numero X: ")
	fmt.Scan(&x)
	fmt.Print("Ingrese el numero Y: ")
	fmt.Scan(&y)

	result = (potencia(x,y)+sumatoria(x)*y)/mcm(x,y)

	fmt.Printf("El resultado del algoritmo dado x=%d e y=%d es de: %d\n", x, y, result)
	
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
	
	// ejercicio09()
	// ejercicio10()
	// ejercicio11()

	// ejercicio08()
}





func dnaSequence(str1 string, str2 string) bool {
	// var sequence []rune = []rune{'A', 'T', 'G', 'C', 'G', 'T', 'A', 'T'}
	strToRune1 := []rune(str1) //  'G', 'C', 'G', 'T', 'A', 'T', 'A', 'T' 
	strToRune2 := []rune(str2) //  'G', 'C', 'G', 'T', 'A', 'T', 'A', 'T' 

	match := false
	
	for i := 0; i < len(strToRune1); i++ {
		
		for j := 0; j < len(strToRune2); j++ {	
			fmt.Println(len(strToRune1))

			if strToRune1[(i+j)%(len(strToRune1))] != strToRune2[(j)] {

				match = false
				break
			}
		}
	}

	return match
}