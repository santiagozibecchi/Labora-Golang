package main

import "fmt"

/*
## Ejercicio 1: Validación de entrada

Escribe un programa en Go que solicite al usuario dos números (numerador y denominador), intente realizar
la división y maneje el caso en el que el denominador sea cero. Si el denominador es cero, imprime un mensaje
de error indicando que la división no es posible. En caso contrario, imprime el resultado de la división con dos decimales.
*/

func division(numerador float64, denominador float64) (float64, error) {
	
	if denominador == 0 {
		return 0, fmt.Errorf("La división no es posible")
	}

	return numerador/denominador, nil
}

func main()  {
	
	var numerador float64
	var denominador float64
	var resultado float64

	fmt.Print("Ingrese el denominador: ")
	fmt.Scan(&denominador)

	fmt.Print("Ingrese el numerador: ")
	fmt.Scan(&numerador)

	resultado, err := division(numerador, denominador)

	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Printf("El resultado de la división es: %.2f\n", resultado)
	}

}