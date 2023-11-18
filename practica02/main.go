package main

import "fmt"

func executeExercise (fn func()) {
	
	fn()

	var optSelected int

	fmt.Println("\n¿Desea seguir realizando el mismo ejercicio? Presione 1 para continuar o 0 para volver al menu")

	_, err := fmt.Scan(&optSelected)
	for optSelected > 1 || optSelected < 0 || err != nil {
		fmt.Println("😡🫵  Por favor! Ingrese 1 para continuar o 0 para volver al menu principal")
		fmt.Println("----- El valor", optSelected, "no es valido -----")
		fmt.Scan(&optSelected)
	}

	for optSelected == 1 {
		fn()
		fmt.Println("\n¿Desea seguir realizando el mismo ejercicio? Presione 1 para continuar o 0 para volver al menu")
		fmt.Scan(&optSelected)
	}

	if optSelected == 0 {
		main()
	}

}

func ejercicio01( numbers []int ) {

	fmt.Println("Debe ingresar 3 numeros para determinar cual es el mayor: ")

	auxNote := numbers[0]

	for _, v := range numbers {
		if (auxNote < v) {
			auxNote = v
		}
	}

	fmt.Println("1. La mayor nota es: ", auxNote)
}

func ejercicio02(naturalNumbers []int){

	firstThreeNumbers := 0
	firstTenNumbers := 0

	for i := 0; i < len(naturalNumbers); i++ {
		if i < 3 {
			firstThreeNumbers+=naturalNumbers[i]
		}

		if i < len(naturalNumbers) {
			firstTenNumbers+=naturalNumbers[i]
		}
	}
	fmt.Println("2. La suma de los 3 primeros numeros naturales es: ", firstThreeNumbers)
	fmt.Println("2. La suma de los 10 primeros numeros naturales es: ", firstTenNumbers)
}

func ejercicio03(){

	var totalNumber int
	var enterNumber int
	var arrayNumber []int
	fmt.Print("Ingrese cuantos numeros quiere sumar: ")
	fmt.Scan(&totalNumber)

	for len(arrayNumber) < totalNumber {
		fmt.Println("Ingrese el numero: ", len(arrayNumber)+1)
		fmt.Scan(&enterNumber)
		arrayNumber = append(arrayNumber, enterNumber)
	}

	var sum int = 0

	for i := 0; i < len(arrayNumber); i++ {
		sum+=arrayNumber[i]
	}

	fmt.Println("La suma de los numeros ingresado es: ", sum)
}

func ejercicio5(){
/*
	Redactar un algoritmo que pida al usuario el ingreso de una serie de números reales e 
	imprima por pantalla el resultado de sumarlos. La serie finaliza cuando el usuario 
	ingresa el número cero. Comparando el ejercicio anterior, qué analogías y diferencias 
	encuentra.
*/

	var enterNumber float64 = 1
	var arrayNumber []float64

	for enterNumber != 0 {
		fmt.Println("Ingrese el numero: ", len(arrayNumber)+1)
		fmt.Scan(&enterNumber)
		arrayNumber = append(arrayNumber, enterNumber)
	}

	var sum float64 = 0

	for i := 0; i < len(arrayNumber); i++ {
		sum+=arrayNumber[i]
	}

	fmt.Println("La suma de los numeros ingresado es: ", sum)

}

func ejercicio06(){
/*
	Ejercicio FizzBuzz en su forma clásica es el siguiente: 
	"Para cada número del 1 al 100, imprime 'Fizz' si el número es divisible por 3, 
	'Buzz' si es divisible por 5 y 'FizzBuzz' si es divisible por ambos. 
	Si no es divisible ni por 3 ni por 5, simplemente imprime el número."
*/
	for i := 0; i <= 100; i++ {
	
		if i % 3 == 0 {
			fmt.Println("Fizz")
		}
		if i % 5 == 0 {
			fmt.Println("Buzz")
		}
		if i % 5 == 0 && i % 3 == 0 {
			fmt.Println("FizzBuzz")
		}
	}
}

func ejercicio07() {

	number1 := 3
	number2 := 7
 
	var divisibleNumbersForN1 []int = []int{}
	var divisibleNumbersForN2 []int = []int{}


	for i := 1; i <= number1; i++ {
		if number1 % i == 0 {
			divisibleNumbersForN1 = append(divisibleNumbersForN1, i)
		}
	}

	for i := 1; i <= number2; i++ {
		if number2 % i == 0 {
			divisibleNumbersForN2 = append(divisibleNumbersForN2, i)
		}
	}

	var tempMCD int = 0
	
	for i := 0; i < len(divisibleNumbersForN1); i++ {
		for j := 0; j < len(divisibleNumbersForN2); j++ {

			if divisibleNumbersForN1[i] == divisibleNumbersForN2[j]{
				tempMCD = divisibleNumbersForN1[i]
			}

		}
	}
	fmt.Printf("El MCD de %d y %d es: %d\n", number1, number2, tempMCD)
}


func ejercicio08(){
/*
	Desarrollar un algoritmo para hallar el mínimo común múltiplo (abreviado como mcm) 
	entre dos números naturales. El mínimo común múltiplo entre dos números es el menor 
	número que es divisible por ambos.
	Ej.: mcm (6,9) = 18, mcm (10,15) = 30, mcm (7,14) = 14, mcm (3,7) = 21
*/

}

func ejercicio09(){
 /*
	Dado un número de 5 cifras, determinar si es capicúa. Si fuera un número de 6 cifras 
	¿Sirve la resolución planteada? ¿Cómo habría que modificarla?
 */

}


func main() {

	// determinar el mayor de 3 números
	// numbers := []int{5,9,7}
	// ejercicio01(numbers)

	// sumar los primeros 3 números naturales. Y luego un algoritmo para sumar los primeros 10 números naturales
	// naturalNumbers := []int{5,3,10,6,25,8,7,9,10,55}
	// ejercicio02(naturalNumbers)

	// ejercicio03()

	// ejercicio5()

	// ejercicio06()

	// ejercicio07()

	// ejercicio08()

	// ejercicio09()


}


